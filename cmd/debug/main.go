package main

import (
	"context"
	"net/http"
	"os"

	"github.com/wt-tools/wtscope/input/hudmsg"
	"github.com/wt-tools/wtscope/net/poll"

	"github.com/grafov/kiwi"
)

func main() {
	ctx := context.Background()
	logger := kiwi.New()
	kiwi.SinkTo(os.Stdout, kiwi.AsLogfmt())
	var localStorage interface{} // XXX
	defaultPolling := poll.New(logger, http.DefaultClient)
	defaultEvents := hudmsg.New(logger, localStorage, defaultPolling)
	defaultEvents.Grab(ctx)
	for {
		ev := defaultEvents.GetDamage(ctx)
		kiwi.Log("event", ev)
	}
}
