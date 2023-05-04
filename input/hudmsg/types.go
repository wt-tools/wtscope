package hudmsg

import (
	"context"
	"encoding/json"
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

type filter interface {
	Important(context.Context) bool
}
type deduplicator interface {
	Exists(uint) bool
}
type configurator interface {
	GamePoint(string) string
}
