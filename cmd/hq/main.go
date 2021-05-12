package main

import (
	"context"
	"net/http"
	"os"

	"github.com/grafov/kiwi"
	"github.com/wt-tools/hq/config"
	"github.com/wt-tools/hq/dedup"
	"github.com/wt-tools/hq/hudmsg"
	"github.com/wt-tools/hq/keep"
	"github.com/wt-tools/hq/poll"
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

	go defaultPolling.Do()
	go hudmsgWorker.Grab(ctx)
	for {
		ev := hudmsgWorker.LatestAction(ctx)
		if ev.Damage != nil {
			if ev.Damage.Player.Name == conf.CurrentPlayer() || ev.Damage.TargetPlayer.Name == conf.CurrentPlayer() {
				log.Log("damage", ev.Origin, "player tank", ev.Damage.Vehicle.Type, "opponent tank", ev.Damage.TargetVehicle.Type, "player", ev.Damage.Player.Name, "target player", ev.Damage.TargetPlayer.Name, "?enemy", ev.Enemy)
			}
		}
	}
}
