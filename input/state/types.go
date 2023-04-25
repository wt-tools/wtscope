package state

import (
	"context"
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

/* 23 april 23, И-180:
{"valid": true,
"aileron, %": 11,
"elevator, %": 21,
"rudder, %": 26,
"flaps, %": 0,
"gear, %": 100,
"H, m": 163,
"TAS, km/h": 0,
"IAS, km/h": 0,
"M": 0.00,
"AoA, deg": -37.0,
"AoS, deg": 22.9,
"Ny": 0.70,
"Vy, m/s": 0.0,
"Wx, deg/s": -0,
"Mfuel, kg": 138,
"Mfuel0, kg": 280,
"throttle 1, %": 0,
"RPM throttle 1, %": 43,
"mixture 1, %": 100,
"radiator 1, %": 0,
"compressor stage 1": 2,
"magneto 1": 3,
"power 1, hp": 12.0,
"RPM 1": 529,
"manifold pressure 1, atm": 0.96,
"oil temp 1, C": 39,
"pitch 1, deg": 25.0,
"thrust 1, kgs": 66,
"efficiency 1, %": 0}
*/

/* sample from april 2023 I-180:
{"valid": true,
"aileron, %": 0,
"elevator, %": -1,
"rudder, %": 0,
"flaps, %": 0,
"gear, %": 100,
"H, m": 175,
"TAS, km/h": 311,
"IAS, km/h": 308,
"M": 0.25,
"AoA, deg": 1.6,
"AoS, deg": 0.0,
"Ny": 1.02,
"Vy, m/s": 3.2,
"Wx, deg/s": 0,
"Mfuel, kg": 138,
"Mfuel0, kg": 280,
"throttle 1, %": 100,
"RPM throttle 1, %": 100,
"mixture 1, %": 100,
"radiator 1, %": 0,
"compressor stage 1": 1,
"magneto 1": 3,
"power 1, hp": 1002.9,
"RPM 1": 2329,
"manifold pressure 1, atm": 1.19,
"oil temp 1, C": 51,
"pitch 1, deg": 29.6,
"thrust 1, kgs": 687,
"efficiency 1, %": 78}
*/

// State keeps original structure as it offered by WT
// `GET state` call.
type state struct {
	Valid                bool    `json:"valid"`
	Aileron              int     `json:"aileron, %"`
	Elevator             int     `json:"elevator, %"`
	Rudder               int     `json:"rudder, %"`
	Flaps                int     `json:"flaps, %"`
	Gear                 int     `json:"gear, %"`
	HM                   int     `json:"H, m"`
	TASKmH               int     `json:"TAS, km/h"`
	IASKmH               int     `json:"IAS, km/h"`
	M                    float64 `json:"M"`
	AoADeg               float64 `json:"AoA, deg"`
	AoSDeg               float64 `json:"AoS, deg"`
	Ny                   float64 `json:"Ny"`
	VyMS                 float64 `json:"Vy, m/s"`
	WxDegS               int     `json:"Wx, deg/s"`
	MfuelKg              int     `json:"Mfuel, kg"`
	Mfuel0Kg             int     `json:"Mfuel0, kg"`
	Throttle1            int     `json:"throttle 1, %"`
	RPMThrottle1         int     `json:"RPM throttle 1, %"`
	Mixture1             int     `json:"mixture 1, %"`
	Radiator1            int     `json:"radiator 1, %"`
	CompressorStage1     int     `json:"compressor stage 1"`
	Magneto1             int     `json:"magneto 1"`
	Power1Hp             float64 `json:"power 1, hp"`
	RPM1                 int     `json:"RPM 1"`
	ManifoldPressure1Atm float64 `json:"manifold pressure 1, atm"`
	OilTemp1C            int     `json:"oil temp 1, C"`
	Pitch1Deg            float64 `json:"pitch 1, deg"`
	Thrust1Kgs           int     `json:"thrust 1, kgs"`
	Efficiency1          int     `json:"efficiency 1, %"`
}

type filter interface {
	Important(context.Context) bool
}
type poller interface {
	Do()
	Add(string, string, int, int) chan []byte
}
type configurator interface {
	GamePoint(string) string
}
