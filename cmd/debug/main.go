package main

import (
	"context"
	"net/http"
	"os"

	"github.com/wt-tools/wtscope/config"
	"github.com/wt-tools/wtscope/input/hudmsg"
	"github.com/wt-tools/wtscope/input/state"
	"github.com/wt-tools/wtscope/net/dedup"
	"github.com/wt-tools/wtscope/net/poll"

	"github.com/grafov/kiwi"
)

func main() {
	ctx := context.Background()
	kiwi.SinkTo(os.Stdout, kiwi.AsLogfmt()).Start()
	log := kiwi.New()
	conf := config.New()
	log.Log("status", "WTScope started")
	defaultPolling := poll.New(http.DefaultClient, nil) // XXX
	hudmsgDedup := dedup.New()
	errlog := make(chan error, 16)
	hudmsgWorker := hudmsg.New(log, conf, defaultPolling, hudmsgDedup)
	stateWorker := state.New(conf, defaultPolling, errlog)
	go defaultPolling.Do()
	go hudmsgWorker.Grab(ctx)
	go stateWorker.Grab(ctx)
	for {
		select {
		case ev := <-hudmsgWorker.Actions(ctx):
			if ev.Damage != nil {
				if ev.Damage.Player.Name == conf.CurrentPlayer() || ev.Damage.TargetPlayer.Name == conf.CurrentPlayer() {
					log.Log("damage", ev.Origin, "player tank", ev.Damage.Vehicle.Name, "opponent tank", ev.Damage.TargetVehicle.Name, "player", ev.Damage.Player.Name, "target player", ev.Damage.TargetPlayer.Name, "?enemy", ev.Enemy)
				}
			}
		case st := <-stateWorker.Actions(ctx):
			log.Log("state", st)

		}
	}
}
