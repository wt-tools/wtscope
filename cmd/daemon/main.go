package main

import (
	"context"
	"net/http"
	"os"

	"github.com/grafov/kiwi"
	"github.com/wt-tools/adjutant/config"
	"github.com/wt-tools/adjutant/dedup"
	"github.com/wt-tools/adjutant/hudmsg"
	"github.com/wt-tools/adjutant/keep"
	"github.com/wt-tools/adjutant/poll"
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
		ev := hudmsgWorker.LatestDamage(ctx)
		log.Log("damage", ev.Origin, "id", ev.ID)
	}
}
