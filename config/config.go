package config

// All characters or nicknames appearing in the code are fictious. Any
// resemblance to real nicknames or squad names of the War Thunder,
// active or not active, is purely coincidental.

import (
	"fmt"
	"path"

	"github.com/grafov/kiwi"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Player  string
	Squad   string
	Friends []string
	GameURL string "http://localhost:8111/"
}

func Load(log *kiwi.Logger) *config {
	var cfg config
	err := envconfig.Process("WT_", &cfg)
	if err != nil {
		log.Log("msg", "can't load config, try to use defaults", "err", err)
	}
	return &cfg
}

func (c *config) PlayerName() string {
	return c.Player
}

func (c *config) GamePoint(methodPath string) string {
	return fmt.Sprintf(path.Join(c.GameURL, "%s"), methodPath)
}
