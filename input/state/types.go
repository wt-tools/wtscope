package state

import (
	"context"

	"github.com/wt-tools/hq/sensor"
)

/* example:
{"valid": true,
"aileron, %": -0,
"elevator, %": -0,
"rudder, %": -0,
"flaps, %": 0,
"gear, %": 100,
"H, m": 59,
"TAS, km/h": 0,
"IAS, km/h": 0,
"M": 0.00,
"AoA, deg": 20.5,
"AoS, deg": 46.3,
"Ny": 0.90,
"Vy, m/s": -0.0,
"Wx, deg/s": -1,
"Mfuel, kg": 91,
"Mfuel0, kg": 303,
"throttle 1, %": 59,
"RPM throttle 1, %": 100,
"radiator 1, %": 0,
"magneto 1": 3,
"power 1, hp": 0.0,
"RPM 1": 0,
"manifold pressure 1, atm": 0.99,
"water temp 1, C": 28,
"oil temp 1, C": 27,
"pitch 1, deg": 24.0,
"thrust 1, kgs": 0,
"efficiency 1, %": 0}
*/

// State keeps original structure as it offered by WT
// `GET state` call.
type state struct {
	Valid        bool    `json:"valid"`
	Aileron      uint8   `json:"aileron, %"`
	Elevator     uint8   `json:"elevator, %"`
	Rudder       uint8   `json:"rudder, %"`
	Flaps        uint8   `json:"flaps, %"`
	Gear         uint8   `json:"gear, %"`
	H            uint    `json:"H, m"`
	TAS          uint    `json:"TAS, km/h"`
	IAS          uint    `json:"IAS, km/h"`
	M            float64 `json:"M"`
	AoA          float64 `json:"AoA, deg"`
	AoS          float64 `json:"AoS, deg"`
	Ny           float64 `json:"Ny"`
	Vy           float64 `json:"Vy"`
	Wx           int8    `json:"Wx, deg/s"`
	Mfuel        int     `json:"Mfuel, kg"`
	Mfuel0       int     `json:"Mfuel0, kg"`
	Throttle1    uint8   `json:"throttle 1, %"`
	Throttle2    uint8   `json:"throttle 2, %"`
	ThrottleRPM1 uint8   `json:"RPM throttle 1, %"`
	ThrottleRPM2 uint8   `json:"RPM throttle 2, %"`
	Radiator1    uint8   `json:"radiator 1, %"`
	Radiator2    uint8   `json:"radiator 2, %"`
	Magneto1     uint8   `json:"magneto 1"`
	Magneto2     uint8   `json:"magneto 2"`
	Power1       float64 `json:"power 1, hp"`
	Power2       float64 `json:"power 2, hp"`
	RPM1         int     `json:"RPM 1"`
	RPM2         int     `json:"RPM 2"`
	Pressure1    float64 `json:"manifold pressure 1, atm"`
	Pressure2    float64 `json:"manifold pressure 2, atm"`
	WaterTemp1   int     `json:"water temp 1, C"`
	WaterTemp2   int     `json:"water temp 2, C"`
	OilTemp1     int     `json:"oil temp 1, C"`
	OilTemp2     int     `json:"oil temp 2, C"`
	Pitch1       float64 `json:"pitch 1, deg"`
	Pitch2       float64 `json:"pitch 2, deg"`
	Thrust1      int     `json:"thrust 1, kgs"`
	Thrust2      int     `json:"thrust 2, kgs"`
	Efficiency1  uint8   `json:"efficiency 1, %"`
	Efficiency2  uint8   `json:"efficiency 2, %"`
}
type keeper interface {
	PersistState(context.Context, sensor.Sensor)
}
type filter interface {
	Important(context.Context) bool
}
type poller interface {
	Add(string, string, int, int) chan []byte
}
type configurator interface {
	GamePoint(string) string
}
