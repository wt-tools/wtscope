package poll

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type httper interface {
	Do(req *http.Request) (*http.Response, error)
}

const RepeatEndlessly = -1

type Service struct {
	sync.Mutex
	queue          []Task
	current        int
	httpc          httper
	loopDelay      time.Duration
	onProblemDelay time.Duration
	err            chan error
}

func New(c httper, logger chan error, opts ...option) *Service {
	s := Service{
		httpc:          c,
		err:            logger,
		loopDelay:      1 * time.Second,
		onProblemDelay: 2 * time.Second,
	}
	for _, o := range opts {
		o(&s)
	}
	return &s
}

// Do tasks syncrhonously.
func (s *Service) Do() {
	var (
		data []byte
		err  error
	)
	for {
		time.Sleep(s.loopDelay)
		s.Lock()
		var t Task
		{
			if len(s.queue) == 0 {
				s.Unlock()
				time.Sleep(s.onProblemDelay)
				continue
			}
			t = s.queue[s.current]
			if s.queue[s.current].repeat != RepeatEndlessly {
				s.queue[s.current].repeat--
			}
			if s.queue[s.current].repeat == 0 {
				close(t.ret)
				s.queue = append(s.queue[:s.current], s.queue[s.current+1:]...)
			}
			s.current++
			if s.current >= len(s.queue) {
				s.current = 0
			}
		}
		s.Unlock()
		if data, err = s.callEndpoint(t); err != nil {
			s.log(err)
			time.Sleep(s.onProblemDelay)
			continue
		}
		t.ret <- data
	}
}

func (s *Service) log(err error) {
	if s.err != nil {
		s.err <- err
	}
}

func (s *Service) callEndpoint(t Task) ([]byte, error) {
	const requestTimeout = 2 * time.Second
	var (
		req *http.Request
		res *http.Response
		err error
	)
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	req, err = http.NewRequestWithContext(ctx, t.method, t.url, nil)
	for r := t.retry; r > 0; r-- {
		if res, err = s.httpc.Do(req); err != nil {
			return nil, err
		}
		if res.StatusCode >= http.StatusBadRequest {
			return nil, fmt.Errorf("bad request (code: %d)", res.StatusCode)
		}
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

// Add request to task queue and repeat it in a loop. After execution
// the task still remains in the queue until `repeat` count decreased
// to 0. For endless repeat set `repeat` to -1. By default no retries
// for the request, set `retry` to value greater than 0.
func (s *Service) Add(method, url string, repeat, retry int) Task {
	if repeat == 0 {
		repeat = 1
	}
	if retry <= 0 {
		retry = 1
	}
	t := Task{method, url, repeat, retry, make(chan []byte, 1)}
	s.Lock()
	s.queue = append(s.queue, t)
	s.Unlock()
	return t
}
