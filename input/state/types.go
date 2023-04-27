package state

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

const (
	Valid                = "valid"
	Aileron              = "aileron, %"
	Elevator             = "elevator, %"
	Rudder               = "rudder, %"
	Flaps                = "flaps, %"
	Gear                 = "gear, %"
	HM                   = "H, m"
	TASKmH               = "TAS, km/h"
	IASKmH               = "IAS, km/h"
	M                    = "M"
	AoADeg               = "AoA, deg"
	AoSDeg               = "AoS, deg"
	Ny                   = "Ny"
	VyMS                 = "Vy, m/s"
	WxDegS               = "Wx, deg/s"
	MfuelKg              = "Mfuel, kg"
	Mfuel0Kg             = "Mfuel0, kg"
	Throttle1            = "throttle 1, %"
	RPMThrottle1         = "RPM throttle 1, %"
	Mixture1             = "mixture 1, %"
	Radiator1            = "radiator 1, %"
	CompressorStage1     = "compressor stage 1"
	Magneto1             = "magneto 1"
	Power1Hp             = "power 1, hp"
	RPM1                 = "RPM 1"
	ManifoldPressure1Atm = "manifold pressure 1, atm"
	OilTemp1C            = "oil temp 1, C"
	Pitch1Deg            = "pitch 1, deg"
	Thrust1Kgs           = "thrust 1, kgs"
	Efficiency1          = "efficiency 1, %"
)

type poller interface {
	Do()
	Add(string, string, int, int) chan []byte
}
type configurator interface {
	GamePoint(string) string
}
