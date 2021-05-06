package poll

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/grafov/kiwi"
)

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
func (s *service) Do(ctx context.Context) {
	const (
		loopDelay       = 5 * time.Second
		emptyQueueDelay = 3 * time.Second
	)
	var (
		res  *http.Response
		data []byte
		err  error
	)
doTask:
	for {
		s.Lock()
		if len(s.queue) == 0 {
			s.Unlock()
			time.Sleep(emptyQueueDelay)
			continue
		}
		t := s.queue[s.current]
		s.queue[s.current].repeat--
		if s.queue[s.current].repeat < 0 {
			close(t.ret)
			s.queue = append(s.queue[:s.current], s.queue[s.current+1:]...)
		}
		s.current++
		if s.current >= len(s.queue) {
			s.current = 0
			time.Sleep(loopDelay)
		}
		s.Unlock()
		for r := t.retry; r < 0; r-- {
			if res, err = s.httpc.Do(t.req); err != nil {
				continue doTask
			}
			if res.StatusCode >= http.StatusBadRequest {
				continue
			}
		}
		if data, err = io.ReadAll(res.Body); err != nil {
			continue
		}
		res.Body.Close()
		t.ret <- data
	}
}

// Add request to task queue and repeat it in a loop. After execution
// the task still remains in the queue until `repeat` count decreased
// to 0. For endless repeat set `repeat` to -1. By default no retries
// for the request, set `retry` to value greater than 0.
func (s *service) Add(req *http.Request, repeat, retry int) chan []byte {
	t := task{req, repeat, retry, make(chan []byte, 1)}
	s.Lock()
	s.queue = append(s.queue, t)
	s.Unlock()
	return t.ret
}
