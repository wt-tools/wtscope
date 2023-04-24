// Package hudmsg parses JSON data input from WT game webserver and
// sends it for further analyzis and to storage.
package hudmsg

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/wt-tools/wtscope/action"
	"github.com/wt-tools/wtscope/net/poll"
	"github.com/wt-tools/wtscope/tag"

	"github.com/grafov/kiwi"
)

type Service struct {
	Messages chan action.GeneralAction

	poll  poller
	filt  filter
	dedup deduplicator
	conf  configurator
	log   *kiwi.Logger
}

func New(log *kiwi.Logger, conf configurator, poll poller, dedup deduplicator) *Service {
	const name = "hudmsg"
	return &Service{
		log:      log.Fork().With(tag.Service, name),
		conf:     conf,
		poll:     poll,
		dedup:    dedup,
		Messages: make(chan action.GeneralAction, 3),
	}
}

func (s *Service) Grab(ctx context.Context) {
	var (
		data []byte
		raw  Raw
		ok   bool
		err  error
	)
	ret := s.poll.Add(http.MethodGet, s.conf.GamePoint("hudmsg&lastEvt=0?lastDmg=0"), poll.RepeatEndlessly, 0)
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
			s.Messages <- dmg
			// s.keep.Cache(ctx, dmg) // XXX
			// if dmg.Important() {
			//	s.keep.Persist(ctx, dmg) // XXX
			// }
		}
	}
}
