package action

type Action int

const (
	Unknown Action = iota
	Connected
	Destroyed
	Got
	Afire
	Wrecked
)

// подбил, сбил, присоединился, потерял связь
