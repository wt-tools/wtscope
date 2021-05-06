Rpackage event

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grafov/kiwi"
	"github.com/wt-tools/adjutant/config"
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
		r      *http.Request
		data   []byte
		rawmsg hudmsg
		damage Damage
		ok     bool
		err    error
	)
	r, err = http.NewRequestWithContext(ctx, http.MethodGet, config.GamePoint("hudmsg"), nil)
	ret := s.poll.Add(r, -1, 0)
	for {
		if data, ok = <-ret; !ok {
			s.log.Log(tag.ExitOn, "channel closed")
			return
		}
		if err = json.Unmarshal(data, &rawmsg); err != nil {
			s.log.Log(tag.Error, err)
			continue
		}
		if damage, err = s.parseDamage(ctx, rawmsg); err != nil {
			s.log.Log()
			continue
		}
		s.filt.Important(ctx)
		s.keep.Save(ctx)
	}
}

func (s *service) parseDamage(ctx context.Context, raw hudmsg) (Damage, error) {
	return Damage{}, nil
}
