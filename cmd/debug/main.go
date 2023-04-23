package main

import (
	"context"
	"net/http"
	"os"
	"wt-tools/wtscope/event"
	"wt-tools/wtscope/poll"

	"github.com/grafov/kiwi"
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
