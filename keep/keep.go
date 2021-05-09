// Package keep allow to keep data in memory (cache) or persistently
// in sqlite.
package keep

import (
	"context"

	"github.com/grafov/kiwi"
	"github.com/wt-tools/adjutant/action"
)

type service struct {
	log *kiwi.Logger
}

// New ...
func New(log *kiwi.Logger) *service {
	return &service{log}
}

func (s *service) Persist(context.Context, action.Damage) {
}

func (s *service) Cache(context.Context, action.Damage) {
}
