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
	logger := kiwi.New()
	kiwi.SinkTo(os.Stdout, kiwi.AsLogfmt())
	localStorage := keep.New(logger)
	defaultPolling := poll.New(logger, http.DefaultClient)
	hudmsgWorker := hudmsg.New(logger, localStorage, defaultPolling)
	go hudmsgWorker.Grab(ctx)
	for {
		ev := hudmsgWorker.LatestDamage(ctx)
		kiwi.Log("event", ev)
	}
}
