package action

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
