// Package state parses JSON data input from WT game webserver and
// sends it for further analyzis and to storage.
package state

import (
	"context"
	"net/http"

	"github.com/valyala/fastjson"
	"github.com/wt-tools/wtscope/net/poll"
)

type Service struct {
	Messages chan *fastjson.Value

	p    fastjson.Parser
	poll poller
	conf configurator
	err  chan error
}

func New(conf configurator, poll poller, log chan error) *Service {
	const name = "state"
	return &Service{
		err:      log,
		conf:     conf,
		Messages: make(chan *fastjson.Value, 3),
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
		var v *fastjson.Value
		if v, err = s.p.ParseBytes(data); err != nil {
			s.log(err)
			continue
		}
		s.Messages <- v
	}
}
