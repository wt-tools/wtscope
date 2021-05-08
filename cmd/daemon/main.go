package main

import (
	"context"
	"net/http"
	"os"

	"github.com/grafov/kiwi"
	"github.com/wt-tools/adjutant/hudmsg"
	"github.com/wt-tools/adjutant/keep"
	"github.com/wt-tools/adjutant/poll"
)

func main() {
	ctx := context.Background()
	kiwi.SinkTo(os.Stdout, kiwi.AsLogfmt()).Start()
	log := kiwi.New()
	log.Log("program", "started")
	localStorage := keep.New(log)
	defaultPolling := poll.New(log, http.DefaultClient)
	hudmsgWorker := hudmsg.New(log, localStorage, defaultPolling)

	go defaultPolling.Do()
	go hudmsgWorker.Grab(ctx)
	for {
		ev := hudmsgWorker.LatestDamage(ctx)
		log.Log("damage", ev.Origin)
	}
}
