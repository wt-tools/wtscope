package hudmsg

import (
	"context"
	"encoding/json"

	"github.com/wt-tools/hq/action"
)

type (
	// GameOutput keeps original structure as it offered by WT
	// `GET hudmsg` call.
	Raw struct {
		Events []Event  `json:"events"`
		Damage []Damage `json:"damage"`
	}
	Event  json.RawMessage // no samples yet
	Damage struct {
		ID     uint   `json:"id"`
		Msg    string `json:"msg"`
		Sender string `json:"sender"`
		Enemy  bool   `json:"enemy"`
		Mode   string `json:"mode"`
	}
)

type keeper interface {
	Persist(context.Context, action.Damage)
	Cache(context.Context, action.Damage)
}
type filter interface {
	Important(context.Context) bool
}
type poller interface {
	Add(string, string, int, int) chan []byte
}
type deduplicator interface {
	Exists(uint) bool
}
type configurator interface {
	GamePoint(string) string
}
