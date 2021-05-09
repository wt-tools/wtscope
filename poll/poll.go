package poll

import (
	"context"
	"errors"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/grafov/kiwi"
	"github.com/wt-tools/adjutant/tag"
)

const RepeatEndlessly = -1

type service struct {
	sync.Mutex
	queue   []task
	current int

	httpc httper
	log   *kiwi.Logger
}

func New(log *kiwi.Logger, c httper) *service {
	return &service{log: log, httpc: c}
}

// Do tasks syncrhonously.
func (s *service) Do() {
	const (
		loopDelay       = 5 * time.Second
		emptyQueueDelay = 3 * time.Second
	)
	var (
		data []byte
		err  error
	)
	for {
		time.Sleep(loopDelay)
		s.Lock()
		var t task
		{
			if len(s.queue) == 0 {
				s.Unlock()
				time.Sleep(emptyQueueDelay)
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
			s.log.Log(tag.Error, err)
			continue
		}
		t.ret <- data
	}
}

func (s *service) callEndpoint(t task) ([]byte, error) {
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
			return nil, errors.New("bad request")
		}
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

// Add request to task queue and repeat it in a loop. After execution
// the task still remains in the queue until `repeat` count decreased
// to 0. For endless repeat set `repeat` to -1. By default no retries
// for the request, set `retry` to value greater than 0.
func (s *service) Add(method, url string, repeat, retry int) chan []byte {
	if repeat == 0 {
		repeat = 1
	}
	if retry <= 0 {
		retry = 1
	}
	t := task{method, url, repeat, retry, make(chan []byte, 1)}
	s.Lock()
	s.queue = append(s.queue, t)
	s.Unlock()
	return t.ret
}
