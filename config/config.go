package config

// All characters or nicknames appearing in the code are fictious. Any
// resemblance to real nicknames or squad names of the War Thunder,
// active or not active, is purely coincidental.

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

const (
	ConfPath     = "~/.config/wtscope/defaults.toml"
	confTemplate = `# Replace with real configuration data
[player]
name = "YourPlayerName"
squad = "YourSquadName"

[client]
url = "http://localhost:8111/"
`
)

type Config struct {
	parsed struct {
		Player    string   `koanf:"player.name"`
		Squad     string   `koanf:"player.squad"`
		Friends   []string `koanf:"player.friends"`
		ClientURL string   `koanf:"client.url"`
	}
	konfig *koanf.Koanf
	err    chan error
}

// Load load configuration. Don't forget to pass log channel.
func Load(log chan error) (*Config, error) {
	var err error
	konfig := koanf.New(".")
	cfg := Config{err: log, konfig: konfig}
	f := file.Provider(preparePath(ConfPath))
	if err = konfig.Load(f, toml.Parser()); err != nil {
		// TODO setup defaults here
		return nil, fmt.Errorf("can't load config: %w", err)
	}
	if err = konfig.UnmarshalWithConf("", &cfg.parsed, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true}); err != nil {
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

func preparePath(p string) string {
	if strings.HasPrefix(p, "~/") {
		u, _ := user.Current()
		p = path.Join(u.HomeDir, p[2:])
	}
	return p
}

// CreateIfAbsent creates a new configuration file based on hardcoded
// template. Use if the config loading failed.
func CreateIfAbsent() error {
	var err error
	p := preparePath(ConfPath)
	if _, err = os.Stat(p); errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(p, []byte(confTemplate), 0o644)
	}
	return err
}

func (c *Config) log(err error) {
	if c.err != nil {
		c.err <- err
	}
}

func (c *Config) Dump() string {
	return c.konfig.Sprint()
}

func (c *Config) PlayerName() string {
	return c.parsed.Player
}

func (c *Config) PlayerSquad() string {
	return c.parsed.Squad
}

func (c *Config) GamePoint(methodPath string) string {
	return fmt.Sprintf(c.parsed.ClientURL+"%s", methodPath)
}
