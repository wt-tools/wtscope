package damage

type Action int

const (
	Unknown Action = iota
	Connected
	Destroyed
	Got
	Fired
	Wrecked
)

// подбил, сбил, присоединился, потерял связь
