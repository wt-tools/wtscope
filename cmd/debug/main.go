package main

import (
	"context"
	"net/http"
	"os"

	"github.com/grafov/kiwi"
	"github.com/wt-tools/hq/event"
	"github.com/wt-tools/hq/poll"
)

func main() {
	ctx := context.Background()
	logger := kiwi.New()
	kiwi.SinkTo(os.Stdout, kiwi.AsLogfmt())
	var localStorage interface{} // XXX
	defaultPolling := poll.New(logger, http.DefaultClient)
	defaultEvents := event.New(logger, localStorage, defaultPolling)
	defaultEvents.Grab(ctx)
	for {
		ev := defaultEvents.GetDamage(ctx)
		kiwi.Log("event", ev)
	}
}
