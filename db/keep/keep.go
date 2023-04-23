// Package keep allow to keep data in memory (cache) or persistently
// in sqlite.
package keep

import (
	"context"

	"github.com/wt-tools/wtscope/action"
	"github.com/wt-tools/wtscope/sensor"

	"github.com/grafov/kiwi"
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

func (s *service) PersistDamage(context.Context, action.Damage) {
}

func (s *service) CacheDamage(context.Context, action.Damage) {
}

// XXX
func (s *service) PersistState(context.Context, sensor.Sensor) {
}
