// Package action represents high level logic for damage&events related data.
// It knows nothing about input formats or storage for the data. It
// isolated from logging and configuration environment.
package action

type (
	Damage struct {
		Action        Action
		Vehicle       Vehicle
		Player        Player
		TargetVehicle Vehicle
		TargetPlayer  Player
	}
	Vehicle struct {
		ID   uint
		Type string
	}
)

// XXX
func (d *Damage) Important() bool {
	// XXX
	return true
}
