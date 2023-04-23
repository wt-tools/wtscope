// Package action represents high level logic for damage&events related data.
// It knows nothing about input formats or storage for the data. It
// isolated from logging and configuration environment.
package action

import "github.com/wt-tools/wtscope/vehicle"

type (
	Damage struct {
		Action        Action
		Vehicle       vehicle.Vehicle
		Player        Player
		TargetVehicle vehicle.Vehicle
		TargetPlayer  Player
	}
)

// XXX
func (d *Damage) Important() bool {
	// XXX
	return true
}
