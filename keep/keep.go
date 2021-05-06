// Package keep allow to keep data in memory (cache) or persistently
// in sqlite.
package keep

import "github.com/grafov/kiwi"

type Service struct {
	log *kiwi.Logger
}

// New ...
func New(log *kiwi.Logger) *Service {
	return &Service{log}
}
