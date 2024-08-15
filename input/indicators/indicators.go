package indicators

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/wt-tools/wtscope/config"
	"github.com/wt-tools/wtscope/net/poll"
)

type Service struct {
	Messages chan indicator

	poll *poll.Service
	conf *config.Config
	err  chan error
}

func New(conf *config.Config, pollsvc *poll.Service, log chan error) *Service {
	return &Service{
		err:      log,
		conf:     conf,
		Messages: make(chan indicator, 3),
		poll:     pollsvc,
	}
}

func (s *Service) log(err error) {
	if s.err != nil {
		s.err <- err
	}
}

func (s *Service) Grab(ctx context.Context) {
	var (
		data []byte
		ind  indicator
		ok   bool
		err  error
	)
	t := s.poll.Add("indicators", http.MethodGet, s.conf.GamePoint("indicators"), "/tmp", poll.RepeatEndlessly, 0)
	for {
		if data, ok = <-t.Results(); !ok {
			s.log(errChanClosed)
			return
		}
		if err = json.Unmarshal(data, &ind); err != nil {
			s.log(err)
			continue
		}
		s.Messages <- ind
	}
}
