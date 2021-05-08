// Package keep allow to keep data in memory (cache) or persistently
// in sqlite.
package keep

import "github.com/grafov/kiwi"

type service struct {
	log *kiwi.Logger
}

// New ...
func New(log *kiwi.Logger) *service {
	return &service{log}
}

func (s *service) Persist() {
}
