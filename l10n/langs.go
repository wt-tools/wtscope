package l10n

// Languages supported by UI in War Thunder.
type Lang string

// TODO return ISO codes with text descriptitons
// type Lang struct {
//	Alpha2 string
//	Text   string
// }

const (
	Auto Lang = ""
	Ru        = "русский"
	En        = "english"
	Cn        = "中国语言"
)
