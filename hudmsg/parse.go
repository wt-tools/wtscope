package hudmsg

import (
	"context"

	"github.com/wt-tools/adjutant/damage"
)

// Just funny to parse it without regexps.
func (s *service) parseDamage(ctx context.Context, msg string) (damage.Damage, error) {
	return damage.Damage{Origin: msg}, nil
	// if who, err = s.parseVehicle(msg); err != nil {
	//	return nil, err
	// }
	// if act, err = s.parseAction(msg); err != nil {
	//	return nil, err
	// }

	// dmg := damage.New(id, who, act, whom)
	// return dmg, nil
}

// func (s *service) parseVehicle(msg string) (damage.Vehicle, error) {
// return nil, nil
// }

// func (s *service) parseAction(msg string) (damage.Action, error) {
// return nil, nil
// }
