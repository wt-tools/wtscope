// Package damage represents high level logic for damage related data.
// It knows nothing about input formats or storage for the data.
// It isolated from logging and configuration environment.
package damage

type (
	Damage struct {
		ID     uint
		Action Action
		Who    Vehicle
		Whom   Vehicle
	}
	Vehicle struct {
		Type     string
		TypeID   uint
		Player   string
		PlayerID uint
	}
)

func New(id uint, who Vehicle, act Action, whom Vehicle) *Damage {
	return &Damage{
		ID:     id,
		Who:    who,
		Action: act,
		Whom:   whom,
	}
}

// XXX
func (d *Damage) Important() bool {
	// XXX
	return true
}
