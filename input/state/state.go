// Package state parses JSON data input from WT game webserver and
// sends it for further analyzis and to storage.
package state

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/wt-tools/wtscope/net/poll"
)

type service struct {
	poll   poller
	conf   configurator
	latest chan state
	err    chan error
}

func New(conf configurator, poll poller, log chan error) *service {
	const name = "state"
	return &service{
		err:    log,
		conf:   conf,
		latest: make(chan state, 3),
		poll:   poll,
	}
}

func (s *service) log(err error) {
	if s.err != nil {
		s.err <- err
	}
}

func (s *service) Actions(ctx context.Context) chan state {
	return s.latest
}

func (s *service) Grab(ctx context.Context) {
	var (
		data []byte
		st   state
		ok   bool
		err  error
	)
	ret := s.poll.Add(http.MethodGet, s.conf.GamePoint("state"), poll.RepeatEndlessly, 0)
	for {
		if data, ok = <-ret; !ok {
			s.log(errChanClosed)
			return
		}
		if err = json.Unmarshal(data, &st); err != nil {
			s.log(err)
			continue
		}
		s.latest <- st
	}
}
