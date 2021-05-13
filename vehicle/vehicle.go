package vehicle

type Vehicle struct {
	ID           uint
	Name         string
	Kind         Kind
	Category     Category
	BR           float32
	Manufacturer Country
	Operator     Country
}

type Kind int

const (
	UnknownKind Kind = iota
	Ground
	Navy
	Aircraft
)

type Category int

const (
	UnknownCat Category = iota
	LightTank
	MediumTank
	HeavTyank
)

type Country int // TODO учесть технику стран, не имеющих своей ветки!
const (
	UnknownCountry Country = iota
	US
	DE
	SU
	EN
	FR
	IT
	CH
	SW
)
