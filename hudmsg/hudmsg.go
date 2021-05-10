// Package hudmsg parses JSON data input from WT game webserver and
// sends it for further analyzis and to storage.
package hudmsg

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grafov/kiwi"
	"github.com/wt-tools/adjutant/action"
	"github.com/wt-tools/adjutant/poll"
	"github.com/wt-tools/adjutant/tag"
)

type service struct {
	keep  keeper
	poll  poller
	filt  filter
	dedup deduplicator
	conf  configurator
	log   *kiwi.Logger
}

func New(log *kiwi.Logger, conf configurator, keep keeper, poll poller, dedup deduplicator) *service {
	const name = "hudmsg"
	return &service{
		log:   log.Fork().With(tag.Service, name),
		conf:  conf,
		keep:  keep,
		poll:  poll,
		dedup: dedup,
	}
}

func (s *service) Grab(ctx context.Context) {
	var (
		data []byte
		raw  Raw
		ok   bool
		err  error
	)
	ret := s.poll.Add(http.MethodGet, s.conf.GamePoint("hudmsg"), poll.RepeatEndlessly, 0)
	for {
		if data, ok = <-ret; !ok {
			s.log.Log(tag.ExitOn, "channel closed")
			return
		}
		if err = json.Unmarshal(data, &raw); err != nil {
			s.log.Log(tag.Error, err)
			continue
		}
		for _, d := range raw.Damage {
			if s.dedup.Exists(d.ID) {
				continue
			}
			dmg, err := parseDamage(d)
			if err != nil {
				s.log.Log(tag.Error, err)
				continue
			}
			latest <- dmg
			// s.keep.Cache(ctx, dmg) // XXX
			// if dmg.Important() {
			//	s.keep.Persist(ctx, dmg) // XXX
			// }
		}
	}
}

var latest = make(chan action.GeneralAction, 3) // XXX

func (s *service) LatestAction(ctx context.Context) action.GeneralAction {
	return <-latest
}
