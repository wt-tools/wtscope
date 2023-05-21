package config

// All characters or nicknames appearing in the code are fictious. Any
// resemblance to real nicknames or squad names of the War Thunder,
// active or not active, is purely coincidental.

import (
	"fmt"
	"path"

	"github.com/grafov/kiwi"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var FilePath = "~/.config/wtscope/defaults.toml"

type config struct {
	Player  string   `koanf:"player.name"`
	Squad   string   `koanf:"player.squad"`
	Friends []string `koanf:"player.friends"`
	GameURL string   `koanf:"game.url"`
	konfig  *koanf.Koanf
}

func Load(log *kiwi.Logger) *config {
	var err error
	l := log.New().With("path", FilePath)
	konfig := koanf.New(".")
	cfg := config{konfig: konfig}
	f := file.Provider(FilePath)
	if err = konfig.Load(f, toml.Parser()); err != nil {
		l.Log("msg", "can't load config, try to use embedded defaults", "err", err)
		// TODO setup defaults here
		return &cfg
	}
	if err = konfig.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true}); err != nil {
		l.Log("msg", "fail configuration file parsing", "err", err)
		return &cfg
	}
	// here code thread unsafe yet
	f.Watch(func(event interface{}, err error) {
		if err != nil {
			l.Log("msg", "configuratiton reloading failed, using old config", "err", err)
			return
		}
		l.Log("msg", "configuration has changed")
		konfig = koanf.New(".")
		cfg.konfig = konfig
		konfig.Load(f, toml.Parser())
		konfig.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true})
	})
	return &cfg
}

func (c *config) Dump() string {
	return c.konfig.Sprint()
}

func (c *config) PlayerName() string {
	return c.Player
}

func (c *config) PlayerSquad() string {
	return c.Squad
}

func (c *config) GamePoint(methodPath string) string {
	return fmt.Sprintf(path.Join(c.GameURL, "%s"), methodPath)
}
