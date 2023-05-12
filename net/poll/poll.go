package poll

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

const RepeatEndlessly = -1

type Service struct {
	sync.Mutex
	pool           []Task
	current        int
	httpc          *http.Client
	loopDelay      time.Duration
	onProblemDelay time.Duration
	err            chan error
}

func New(opts ...option) *Service {
	s := Service{
		httpc:          http.DefaultClient,
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
			if len(s.pool) == 0 {
				s.Unlock()
				time.Sleep(s.onProblemDelay)
				continue
			}
			t = s.pool[s.current]
			if s.pool[s.current].repeat != RepeatEndlessly {
				s.pool[s.current].repeat--
			}
			if s.pool[s.current].repeat == 0 {
				close(t.ret)
				s.pool = append(s.pool[:s.current], s.pool[s.current+1:]...)
			}
			s.current++
			if s.current >= len(s.pool) {
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
		if t.log != nil {
			t.log.Write(bytes.Replace(data, []byte("\n"), []byte(" "), -1))
			t.log.Write([]byte("\n"))
		}
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
func (s *Service) Add(name, method, url string, logPath string, repeat, retry int) Task {
	if repeat == 0 {
		repeat = 1
	}
	if retry <= 0 {
		retry = 1
	}
	var (
		f   *os.File
		err error
	)
	if logPath != "" {
		fname := path.Join(logPath, fmt.Sprintf("%s-%s.log", name, time.Now().Format("06-01-02_15:04:05")))
		if f, err = os.Create(fname); err != nil {
			s.log(fmt.Errorf("for %s can't create %s: %w", name, fname, err))
		}
	}
	t := Task{name, method, url, repeat, retry, f, make(chan []byte, 1)}
	s.Lock()
	s.pool = append(s.pool, t)
	s.Unlock()
	return t
}
