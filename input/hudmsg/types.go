package hudmsg

import (
	"encoding/json"

	"github.com/wt-tools/wtscope/net/poll"
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
		Time   uint   `json:"time"` // probably time in seconds from start of the battle?
	}
)

type poller interface {
	Do()
	Add(name, method, url string, logPath string, repeat, retry int) poll.Task
}

type deduplicator interface {
	Exists(uint) bool
	BlockContent([]byte) bool
}
type configurator interface {
	GamePoint(string) string
}
