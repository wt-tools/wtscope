package indicators

/* Reference data:
{"valid": true,
"type": "i_180",
"speed": -0.004712,
"pedals1": 0.316330,
"pedals2": 0.316423,
"pedals3": 0.316500,
"pedals4": 0.316564,
"pedals5": 0.316618,
"pedals6": 0.316663,
"stick_elevator": 0.261600,
"stick_ailerons": 0.163051,
"vario": 0.003464,
"altitude_hour": 173.732040,
"altitude_min": 173.732040,
"bank": 0.023323,
"turn": 0.000143,
"compass": 52.139954,
"clock_hour": 11.033334,
"clock_min": 2.000000,
"clock_sec": 5.000000,
"manifold_pressure": 0.959291,
"rpm": 528.992249,
"oil_pressure": 38.743378,
"oil_pressure1": -273.149994,
"oil_temperature": 38.743378,
"head_temperature": 104.933624,
"mixture": 0.833333,
"mixture_1": 0.833333,
"fuel": 137.749924,
"fuel_pressure": 0.000000,
"gears": 1.000000,
"gear_lamp_down": 1.000000,
"gear_lamp_up": 1.000000,
"gear_lamp_off": 1.000000,
"flaps": 0.000000,
"throttle": 0.000000,
"throttle_1": 0.000000,
"weapon1": 0.000000,
"weapon2": 0.000000,
"weapon3": 0.000000,
"prop_pitch": 0.425532,
"supercharger": 0.999999,
"flaps_indicator": 0.000000,
"gear_l_indicator": 1.000000,
"gear_r_indicator": 1.000000,
"radiator_lever1_1": 0.000000,
"radiator_lever1_2": 0.000000,
"oil_radiator_lever1_1": 0.000000,
"oil_radiator_lever1_2": 0.000000,
"blister1": 1.000000,
"blister2": 1.000000,
"blister3": 1.000000,
"blister4": 1.000000,
"blister5": 1.000000,
"blister6": 1.000000}
*/

type indicator struct {
	AltitudeHour       float64 `json:"altitude_hour"`
	AltitudeMin        float64 `json:"altitude_min"`
	AmmoCounter1       float64 `json:"ammo_counter1"`
	AmmoCounter1Lamp   float64 `json:"ammo_counter1_lamp"`
	AmmoCounter2       float64 `json:"ammo_counter2"`
	AmmoCounter2Lamp   float64 `json:"ammo_counter2_lamp"`
	AmmoCounter3       float64 `json:"ammo_counter3"`
	AmmoCounter3Lamp   float64 `json:"ammo_counter3_lamp"`
	AviahorizonPitch   float64 `json:"aviahorizon_pitch"`
	AviahorizonRoll    float64 `json:"aviahorizon_roll"`
	Bank               float64 `json:"bank"`
	Blister1           float64 `json:"blister1"`
	Blister2           float64 `json:"blister2"`
	Blister3           float64 `json:"blister3"`
	Blister4           float64 `json:"blister4"`
	Blister5           float64 `json:"blister5"`
	Blister6           float64 `json:"blister6"`
	ClockHour          float64 `json:"clock_hour"`
	ClockMin           float64 `json:"clock_min"`
	ClockSec           float64 `json:"clock_sec"`
	Compass            float64 `json:"compass"`
	Flaps              float64 `json:"flaps"`
	FlapsIndicator     float64 `json:"flaps_indicator"`
	Fuel               float64 `json:"fuel"`
	Fuel1              float64 `json:"fuel1"`
	FuelPressure       float64 `json:"fuel_pressure"`
	GearLIndicator     float64 `json:"gear_l_indicator"`
	GearLampDown       float64 `json:"gear_lamp_down"`
	GearLampOff        float64 `json:"gear_lamp_off"`
	GearLampUp         float64 `json:"gear_lamp_up"`
	GearRIndicator     float64 `json:"gear_r_indicator"`
	Gears              float64 `json:"gears"`
	HeadTemperature    float64 `json:"head_temperature"`
	ManifoldPressure   float64 `json:"manifold_pressure"`
	Mixture            float64 `json:"mixture"`
	Mixture1           float64 `json:"mixture_1"`
	OilPressure        float64 `json:"oil_pressure"`
	OilPressure1       float64 `json:"oil_pressure1"`
	OilRadiatorLever11 float64 `json:"oil_radiator_lever1_1"`
	OilRadiatorLever12 float64 `json:"oil_radiator_lever1_2"`
	OilTemperature     float64 `json:"oil_temperature"`
	Pedals1            float64 `json:"pedals1"`
	Pedals2            float64 `json:"pedals2"`
	Pedals3            float64 `json:"pedals3"`
	Pedals4            float64 `json:"pedals4"`
	Pedals5            float64 `json:"pedals5"`
	Pedals6            float64 `json:"pedals6"`
	PropPitch          float64 `json:"prop_pitch"`
	PropPitchHour      float64 `json:"prop_pitch_hour"`
	PropPitchMin       float64 `json:"prop_pitch_min"`
	RadiatorLever11    float64 `json:"radiator_lever1_1"`
	RadiatorLever12    float64 `json:"radiator_lever1_2"`
	Rpm                float64 `json:"rpm"`
	Speed              float64 `json:"speed"`
	StickAilerons      float64 `json:"stick_ailerons"`
	StickElevator      float64 `json:"stick_elevator"`
	Supercharger       float64 `json:"supercharger"`
	Throttle           float64 `json:"throttle"`
	Throttle1          float64 `json:"throttle_1"`
	Trimmer            float64 `json:"trimmer"`
	Turn               float64 `json:"turn"`
	Type               string  `json:"type"`
	Valid              bool    `json:"valid"`
	Vario              float64 `json:"vario"`
	WaterTemperature   float64 `json:"water_temperature"`
	Weapon1            float64 `json:"weapon1"`
	Weapon2            float64 `json:"weapon2"`
	Weapon3            float64 `json:"weapon3"`
}

type poller interface {
	Add(string, string, int, int) chan []byte
}
type deduplicator interface {
	Exists(uint) bool
}
type configurator interface {
	GamePoint(string) string
}
