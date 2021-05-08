package damage

type Action int

const (
	Unknown Action = iota
	Destroyed
	Gived
	Fired
	Wrecked
)
