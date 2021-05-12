// Package keep allow to keep data in memory (cache) or persistently
// in sqlite.
package keep

import (
	"context"

	"github.com/grafov/kiwi"
	"github.com/wt-tools/hq/action"
)

type service struct {
	log *kiwi.Logger
}

type driver interface {
	Init(sql []string)
	Save() // XXX
	Load() // XXX
}

// New ...
func New(log *kiwi.Logger) *service {
	return &service{log}
}

func (s *service) Persist(context.Context, action.Damage) {
}

func (s *service) Cache(context.Context, action.Damage) {
}
