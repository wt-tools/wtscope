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
	const name = "event"
	return &service{keep: keep, log: log.Fork().With(tag.Service, name)}
}

func (s *service) Grab(ctx context.Context) {
	var (
		r    *http.Request
		data []byte
		raw  Raw
		ok   bool
		err  error
	)
	r, err = http.NewRequestWithContext(ctx, http.MethodGet, config.GamePoint("hudmsg"), nil)
	ret := s.poll.Add(r, -1, 0)
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
			dmg, err := s.parseDamage(ctx, d.Msg)
			if err != nil {
				s.log.Log(tag.Error, err)
				break
			}
			s.keep.Cache(ctx, dmg) // XXX
			if dmg.Important() {
				s.keep.Persist(ctx, dmg) // XXX
			}
		}
	}
}

func (s *service) LatestDamage(ctx context.Context) damage.Damage {
	var dmg damage.Damage

	return dmg
}
