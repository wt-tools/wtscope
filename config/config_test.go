package config

import (
	"testing"
)

func TestConfigLoading(t *testing.T) {
	FilePath = "./config_example.toml"
	cfg, err := Load(nil)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	if cfg.Player != "YourPlayerName" {
		t.Log(cfg)
		t.Fail()
	}
	if cfg.Squad != "YourSquadName" {
		t.Log(cfg)
		t.Fail()
	}
}
