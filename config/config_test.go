package config

import (
	"os"
	"testing"

	"github.com/grafov/kiwi"
)

func init() {
	kiwi.SinkTo(os.Stdout, kiwi.AsLogfmt()).Start()
}

var log = kiwi.New()

func TestConfigLoading(t *testing.T) {
	FilePath = "./config_example.toml"
	cfg := Load(log)
	if cfg.Player != "PlayerName" {
		t.Log(cfg)
		t.Fail()
	}
}
