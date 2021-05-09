package hudmsg

import (
	"context"

	"github.com/wt-tools/adjutant/action"
)

// Just funny to parse it without regexps.
func (s *service) parseDamage(ctx context.Context, msg string, id uint) (action.Damage, error) {
	return action.Damage{Origin: msg, ID: id}, nil
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
