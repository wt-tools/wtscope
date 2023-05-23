package main

import (
	"context"
	"net/http"
	"os"

	"github.com/wt-tools/wtscope/config"
	"github.com/wt-tools/wtscope/input/hudmsg"
	"github.com/wt-tools/wtscope/input/indicators"
	"github.com/wt-tools/wtscope/input/state"
	"github.com/wt-tools/wtscope/net/dedup"
	"github.com/wt-tools/wtscope/net/poll"

	"github.com/grafov/kiwi"
)

// example is broken
func main() {
	ctx := context.Background()
	kiwi.SinkTo(os.Stdout, kiwi.AsLogfmt()).Start()
	log := kiwi.New()
	conf := config.New()
	log.Log("status", "WTScope started")
	defaultPolling := poll.New(http.DefaultClient, nil, 3, 2) // XXX
	hudmsgDedup := dedup.New()
	errlog := make(chan error, 16)
	hudmsgSvc := hudmsg.New(log, conf, defaultPolling, hudmsgDedup)
	stateSvc := state.New(conf, defaultPolling, errlog)
	indicatorsSvc := indicators.New(conf, defaultPolling, errlog)
	go defaultPolling.Do()
	go hudmsgSvc.Grab(ctx)
	go stateSvc.Grab(ctx)
	go indicatorsSvc.Grab(ctx)
	for {
		select {
		case ev := <-hudmsgSvc.Messages:
			if ev.Damage != nil {
				if ev.Damage.Player.Name == conf.CurrentPlayer() || ev.Damage.TargetPlayer.Name == conf.CurrentPlayer() {
					log.Log("damage", ev.Origin, "player tank", ev.Damage.Vehicle.Name, "opponent tank", ev.Damage.TargetVehicle.Name, "player", ev.Damage.Player.Name, "target player", ev.Damage.TargetPlayer.Name, "?enemy", ev.Enemy)
				}
			}
		case data := <-stateSvc.Messages:
			log.Log("state", data)
		case data := <-indicatorsSvc.Messages:
			log.Log("indicator", data)

		}
	}
}
