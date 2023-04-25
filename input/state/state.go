// Package state parses JSON data input from WT game webserver and
// sends it for further analyzis and to storage.
package state

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/wt-tools/wtscope/net/poll"
)

type Service struct {
	Messages chan state

	poll poller
	conf configurator
	err  chan error
}

func New(conf configurator, poll poller, log chan error) *Service {
	const name = "state"
	return &Service{
		err:      log,
		conf:     conf,
		Messages: make(chan state, 3),
		poll:     poll,
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
		ok   bool
		err  error
	)
	ret := s.poll.Add(http.MethodGet, s.conf.GamePoint("state"), poll.RepeatEndlessly, 0)
	for {
		if data, ok = <-ret; !ok {
			s.log(errChanClosed)
			return
		}
		// The state returned by WT has tags with commas.
		// This is the problem:
		// https://github.com/golang/go/issues/15000 Packages
		// like easyjson also can't do it. If you know the
		// package that supports JSON keys with commas, let me
		// know.
		m := map[string]interface{}{}
		if err = json.Unmarshal(data, &m); err != nil {
			s.log(err)
			continue
		}
		s.Messages <- mapkeys(m)
	}
}
