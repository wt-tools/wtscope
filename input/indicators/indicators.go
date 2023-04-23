package indicators

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/wt-tools/wtscope/net/poll"
)

type service struct {
	Messages chan indicator

	poll poller
	conf configurator
	err  chan error
}

func New(conf configurator, poll poller, log chan error) *service {
	return &service{
		err:      log,
		conf:     conf,
		Messages: make(chan indicator, 3),
		poll:     poll,
	}
}

func (s *service) log(err error) {
	if s.err != nil {
		s.err <- err
	}
}

func (s *service) Grab(ctx context.Context) {
	var (
		data []byte
		ind  indicator
		ok   bool
		err  error
	)
	ret := s.poll.Add(http.MethodGet, s.conf.GamePoint("indicators"), poll.RepeatEndlessly, 0)
	for {
		if data, ok = <-ret; !ok {
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