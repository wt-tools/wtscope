package state

import (
	"encoding/json"
	"testing"
)

func TestJsonInput(t *testing.T) {
	// sample from april 2023 I-180:
	const i = `
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
`
	var o state
	err := json.Unmarshal([]byte(i), &o)
	if err != nil {
		t.Log(err)
		t.Fatal()
		return
	}
	if o.IASKmH != 308 {
		t.Log("expected IAS = 308 got ", o.IASKmH)
	}
	if o.TASKmH != 311 {
		t.Log("expected IAS = 311 got ", o.IASKmH)
	}
}
