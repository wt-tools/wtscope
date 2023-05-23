package action

//go:generate stringer -type=Code

type Code int

const (
	Unknown Code = iota
	Connected
	Disconnected
	NetworkDisconnect
	Damaged
	Destroyed
	ShotDown
	Achieved
	Afire
	Wrecked
	Crashed
	SoftLanding
	FinalBlow
)
