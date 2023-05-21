package config

// All characters or nicknames appearing in the code are fictious. Any
// resemblance to real nicknames or squad names of the War Thunder,
// active or not active, is purely coincidental.

import (
	"fmt"
	"os/user"
	"path"
	"strings"

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
	err     chan error
}

func Load(log chan error) (*config, error) {
	var err error
	konfig := koanf.New(".")
	cfg := config{err: log, konfig: konfig}
	if strings.HasPrefix(FilePath, "~/") {
		u, _ := user.Current()
		FilePath = path.Join(u.HomeDir, FilePath[2:])
	}
	f := file.Provider(FilePath)
	if err = konfig.Load(f, toml.Parser()); err != nil {
		// TODO setup defaults here
		return nil, fmt.Errorf("can't load config: %w", err)
	}
	if err = konfig.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true}); err != nil {
		return &cfg, fmt.Errorf("fail configuration file parsing: %w", err)
	}
	// here code thread unsafe yet
	f.Watch(func(event interface{}, err error) {
		if err != nil {
			cfg.log(err)
			return
		}
		konfig = koanf.New(".")
		cfg.konfig = konfig
		konfig.Load(f, toml.Parser())
		if err = konfig.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true}); err != nil {
			cfg.log(err)
		}
	})
	return &cfg, nil
}

func (c *config) log(err error) {
	if c.err != nil {
		c.err <- err
	}
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
