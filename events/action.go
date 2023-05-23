// Package action represents high level logic for damage&events related data.
// It knows nothing about input formats or storage for the data. It
// isolated from logging and configuration environment.
package events

import (
	"time"

	"github.com/wt-tools/wtscope/action"
	"github.com/wt-tools/wtscope/l10n"
	"github.com/wt-tools/wtscope/vehicle"
)

type (
	// Repesents semantically parsed action from the log.
	Event struct {
		ID            uint
		Action        action.Code
		ActionText    string    // as it occured in log, not translated
		Lang          l10n.Lang // language of the log message
		Vehicle       vehicle.Vehicle
		Player        Player
		TargetVehicle vehicle.Vehicle
		TargetPlayer  Player
		Achievement   *Achievement
		Origin        string
		Enemy         bool
		At            time.Duration
	}
	// Respesents info about player and the squadron.
	Player struct {
		ID    uint
		Squad string
		Name  string
	}
)
