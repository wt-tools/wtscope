// Sensor is domain model representing state of currently running
// vehicle.
package sensor

import "time"

type Sensor struct {
	At          time.Time
	Description string
}

type Flaps struct {
	Sensor
	Value uint8
}

type Throttle struct {
	Sensor
	Index uint8
	Value uint8
}

type Power struct {
	Sensor
	Index uint8
	Value float64
}

// func RegisterPower() Power {

// }
