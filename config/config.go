package config

// All characters or nicknames appearing in the code are fictious. Any
// resemblance to real nicknames or squad names of the War Thunder,
// active or not active, is purely coincidental.

import "fmt"

type local struct{}

func New() *local {
	return &local{}
}

func (l *local) CurrentPlayer() string {
	return "ZenAviator" // my nickname in game, for development
}

func (l *local) GamePoint(path string) string {
	// XXX
	return fmt.Sprintf("http://localhost:9222/%s", path)
	// return fmt.Sprintf("http://localhost:8111/%s?lastEvt=0&lastDmg=10", path)
}
