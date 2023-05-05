// Package hudmsg parses JSON data input from WT game webserver and
// sends it for further analyzis and to storage.
package hudmsg

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wt-tools/wtscope/action"
	"github.com/wt-tools/wtscope/net/poll"
)

type Service struct {
	Messages chan action.GeneralAction

	poll  *poll.Service
	dmgID uint
	evtID uint
	filt  filter
	dedup deduplicator
	conf  configurator
	err   chan error
}

func New(conf configurator, pollsvc *poll.Service, dedup deduplicator, log chan error) *Service {
	const name = "hudmsg"
	return &Service{
		err:      log,
		conf:     conf,
		poll:     pollsvc,
		dedup:    dedup,
		Messages: make(chan action.GeneralAction, 3),
	}
}

func (s *Service) log(err error) {
	if s.err != nil {
		s.err <- err
	}
}

func (s *Service) hudURL(dmg uint) string {
	s.dmgID = dmg
	return fmt.Sprintf("hudmsg&lastEvt=%d?lastDmg=%d", s.evtID, s.dmgID)
}

func (s *Service) Grab(ctx context.Context) {
	var (
		data []byte
		raw  Raw
		ok   bool
		err  error
	)
	t := s.poll.Add(http.MethodGet, s.conf.GamePoint(s.hudURL(0)), poll.RepeatEndlessly, 0)
	for {
		if data, ok = <-t.Results(); !ok {
			s.log(errChanClosed)
			return
		}

		if err = json.Unmarshal(data, &raw); err != nil {
			s.log(err)
			continue
		}
		for _, d := range raw.Damage {
			if s.dedup.Exists(d.ID) {
				continue
			}
			if s.dmgID != d.ID {
				t.Update(s.hudURL(d.ID))
			}
			dmg, err := parseDamage(d)
			if err != nil {
				s.log(fmt.Errorf("req: %s, err: %w", s.hudURL(d.ID), err))
				continue
			}
			s.Messages <- dmg
		}
	}
}
