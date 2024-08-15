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
	Valid                = "valid"                    // bool
	Aileron              = "aileron, %"               // int
	Elevator             = "elevator, %"              // int
	Rudder               = "rudder, %"                // int
	Flaps                = "flaps, %"                 // int
	Gear                 = "gear, %"                  // int
	HM                   = "H, m"                     // int
	TASKmH               = "TAS, km/h"                // int
	IASKmH               = "IAS, km/h"                // int
	M                    = "M"                        // float64
	AoADeg               = "AoA, deg"                 // float64
	AoSDeg               = "AoS, deg"                 // float64
	Ny                   = "Ny"                       // float64
	VyMS                 = "Vy, m/s"                  // float64
	WxDegS               = "Wx, deg/s"                // int
	MfuelKg              = "Mfuel, kg"                // int
	Mfuel0Kg             = "Mfuel0, kg"               // int
	Throttle1            = "throttle 1, %"            // int
	RPMThrottle1         = "RPM throttle 1, %"        // int
	Mixture1             = "mixture 1, %"             // int
	Radiator1            = "radiator 1, %"            // int
	CompressorStage1     = "compressor stage 1"       // int
	Magneto1             = "magneto 1"                // int
	Power1Hp             = "power 1, hp"              // float64
	RPM1                 = "RPM 1"                    // int
	ManifoldPressure1Atm = "manifold pressure 1, atm" // float64
	OilTemp1C            = "oil temp 1, C"            // int
	Pitch1Deg            = "pitch 1, deg"             // float64
	Thrust1Kgs           = "thrust 1, kgs"            // int
	Efficiency1          = "efficiency 1, %"          // int
)
