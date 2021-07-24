// Package state parses JSON data input from WT game webserver and
// sends it for further analyzis and to storage.
package state

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grafov/kiwi"
	"github.com/wt-tools/hq/net/poll"
	"github.com/wt-tools/hq/tag"
)

type service struct {
	keep keeper
	poll poller
	conf configurator
	log  *kiwi.Logger
}

func New(log *kiwi.Logger, conf configurator, keep keeper, poll poller) *service {
	const name = "state"
	return &service{
		log:  log.Fork().With(tag.Service, name),
		conf: conf,
		keep: keep,
		poll: poll,
	}
}

var latest = make(chan state, 3) // XXX

func (s *service) Get(ctx context.Context) chan state {

}

func (s *service) Grab(ctx context.Context) {
	var (
		data  []byte
		state state
		ok    bool
		err   error
	)
	ret := s.poll.Add(http.MethodGet, s.conf.GamePoint("state"), poll.RepeatEndlessly, 0)
	for {
		if data, ok = <-ret; !ok {
			s.log.Log(tag.ExitOn, "channel closed")
			return
		}
		if err = json.Unmarshal(data, &state); err != nil {
			s.log.Log(tag.Error, err)
			continue
		}
		latest <- state // XXX
	}
}
