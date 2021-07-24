package main

import (
	"context"
	"net/http"
	"os"

	"github.com/grafov/kiwi"
	"github.com/wt-tools/hq/config"
	"github.com/wt-tools/hq/db/keep"
	"github.com/wt-tools/hq/dedup"
	"github.com/wt-tools/hq/input/hudmsg"
	"github.com/wt-tools/hq/input/state"
	"github.com/wt-tools/hq/net/poll"
)

func main() {
	ctx := context.Background()
	kiwi.SinkTo(os.Stdout, kiwi.AsLogfmt()).Start()
	log := kiwi.New()
	log.Log("status", "adjutant at your service", "config", "xxx")
	conf := config.New()
	localStorage := keep.New(log)
	defaultPolling := poll.New(log, http.DefaultClient)
	hudmsgDedup := dedup.New(log)
	hudmsgWorker := hudmsg.New(log, conf, localStorage, defaultPolling, hudmsgDedup)
	stateWorker := state.New(log, conf, localStorage, defaultPolling)
	go defaultPolling.Do()
	go hudmsgWorker.Grab(ctx)
	go stateWorker.Grab(ctx)
	for {
		select {
		case ev := hudmsgWorker.LatestAction(ctx):
			if ev.Damage != nil {
				if ev.Damage.Player.Name == conf.CurrentPlayer() || ev.Damage.TargetPlayer.Name == conf.CurrentPlayer() {
					log.Log("damage", ev.Origin, "player tank", ev.Damage.Vehicle.Name, "opponent tank", ev.Damage.TargetVehicle.Name, "player", ev.Damage.Player.Name, "target player", ev.Damage.TargetPlayer.Name, "?enemy", ev.Enemy)
				}
			}
		case st := stateWorker.LatestState(ctx):
			log.Log("state", st)

		}
	}
}
