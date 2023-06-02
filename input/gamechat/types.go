package gamechat

import (
	"github.com/wt-tools/wtscope/net/poll"
)

type (
	Chat    []Message
	Message struct {
		ID     uint   `json:"id"`
		Msg    string `json:"msg"`
		Sender string `json:"sender"`
		Enemy  bool   `json:"enemy"`
		Mode   string `json:"mode"`
		Time   uint   `json:"time"`
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
type Config interface {
	GamePoint(string) string
}
