// Package hudmsg parses JSON data input from WT game webserver and
// sends it for further analyzis and to storage.
package hudmsg

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grafov/kiwi"
	"github.com/wt-tools/adjutant/config"
	"github.com/wt-tools/adjutant/damage"
	"github.com/wt-tools/adjutant/poll"
	"github.com/wt-tools/adjutant/tag"
)

type service struct {
	keep keeper
	poll poller
	filt filter
	conf configurator
	log  *kiwi.Logger
}

func New(log *kiwi.Logger, keep keeper, poll poller) *service {
	const name = "hudmsg"
	return &service{
		log:  log.Fork().With(tag.Service, name),
		keep: keep,
		poll: poll,
	}
}

func (s *service) Grab(ctx context.Context) {
	var (
		data []byte
		raw  Raw
		ok   bool
		err  error
	)
	ret := s.poll.Add(http.MethodGet, config.GamePoint("hudmsg"), poll.RepeatEndlessly, 0)
	for {
		if data, ok = <-ret; !ok {
			s.log.Log(tag.ExitOn, "channel closed")
			return
		}
		s.log.Log("message get", data)
		if err = json.Unmarshal(data, &raw); err != nil {
			s.log.Log(tag.Error, err)
			continue
		}
		for _, d := range raw.Damage {
			dmg, err := s.parseDamage(ctx, d.Msg)
			if err != nil {
				s.log.Log(tag.Error, err)
				break
			}
			latest <- dmg
			// s.keep.Cache(ctx, dmg) // XXX
			// if dmg.Important() {
			//	s.keep.Persist(ctx, dmg) // XXX
			// }
		}
	}
}

var latest = make(chan damage.Damage, 3) // XXX

func (s *service) LatestDamage(ctx context.Context) damage.Damage {
	return <-latest
}
