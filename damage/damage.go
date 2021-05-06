// Package damage represents high level logic for damage related data.
// It knows nothing about input formats or storage for the data.
// It isolated from logging and configuration environment.
package damage

import (
	"time"
)

type (
	Damage struct {
		Action   string
		ActionID ActionID
		Who      Vehicle
		Whom     Vehicle
		At       time.Time
	}
	Vehicle struct {
		Type     string
		TypeID   uint
		Player   string
		PlayerID uint
	}
)

func New(act Action) *Damage {

}

func (d *Damage) Important() bool {

}
