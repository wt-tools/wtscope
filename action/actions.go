package action

import "time"

type Action int

const (
	Unknown Action = iota
	Connected
	LostConnect
	Disconnected
	Damaged
	Destroyed
	ShotDown
	Achieved
	Afire
	Wrecked
	SoftLanding
	FinalBlow
)

type (
	GeneralAction struct {
		ID          uint
		Damage      *Damage
		Achievement *Achievement
		Origin      string
		Enemy       bool
		At          time.Duration
	}
	Player struct {
		ID   uint
		Clan string
		Name string
	}
)
