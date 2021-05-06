package hudmsg

import (
	"context"
	"encoding/json"
	"net/http"
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
	Save(context.Context)
}
type filter interface {
	Important(context.Context) bool
}
type poller interface {
	Add(*http.Request, int, int) chan []byte
}
type configurator interface {
	Username() string
}
