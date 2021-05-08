package config

import "fmt"

func GamePoint(path string) string {
	// XXX
	return fmt.Sprintf("http://localhost:8111/%s?lastEvt=0&lastDmg=700", path)
}
