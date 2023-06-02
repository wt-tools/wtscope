// Package hudmsg parses JSON data input from WT game webserver and
// sends it for further analyzis and to storage.
package gamechat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/wt-tools/wtscope/net/poll"
)

type Service struct {
	Messages chan Message

	poll  poller
	msgID uint
	dedup deduplicator
	conf  Config
	err   chan error
}

func New(conf Config, pollsvc poller, dedup deduplicator, log chan error) *Service {
	const name = "gamechat"
	return &Service{
		err:      log,
		conf:     conf,
		poll:     pollsvc,
		dedup:    dedup,
		Messages: make(chan Message, 6),
	}
}

func (s *Service) log(err error) {
	if s.err != nil {
		s.err <- err
	}
}

func (s *Service) chatURL(id uint) string {
	s.msgID = id
	return fmt.Sprintf("gamechat?lastId=%d", s.msgID)
}

func (s *Service) Grab(ctx context.Context) {
	var (
		data []byte
		ok   bool
		err  error
	)
	t := s.poll.Add("gamechat", http.MethodGet, s.conf.GamePoint(s.chatURL(0)), "/tmp", poll.RepeatEndlessly, 0)
	for {
		if data, ok = <-t.Results(); !ok {
			s.log(errChanClosed)
			return
		}
		var chat Chat
		if err = json.Unmarshal(data, &chat); err != nil {
			s.log(err)
			continue
		}
		for _, m := range chat {
			if s.dedup.Exists(m.ID) {
				continue
			}
			if s.dedup.BlockContent([]byte(m.Msg + m.Sender + m.Mode)) {
				continue
			}
			if s.msgID != m.ID {
				t.Update(s.chatURL(m.ID))
			}
			m.At = time.Duration(m.Time) * time.Second
			s.Messages <- m
		}
	}
}
