// Package locale defines index for tokens alongside with their
// translations.
package locale

import "github.com/wt-tools/wtscope/action"

// import "github.com/wt-tools/wtscope/action"

// broken
type Translation struct {
	Lang  Lang
	Index action.Damage // broken
	Value string
}

var (
	translationByIndex = make(map[Lang]map[action.Damage]string)
	translationByText  = make(map[Lang]map[string]action.Damage)
)

func init() {
	for _, v := range actionTexts {
		if _, ok := translationByIndex[v.Lang]; !ok {
			translationByIndex[v.Lang] = make(map[action.Damage]string)
		}
		translationByIndex[v.Lang][v.Index] = v.Value
		if _, ok := translationByText[v.Lang]; !ok {
			translationByIndex[v.Lang] = make(map[action.Damage]string)
		}
		translationByText[v.Lang][v.Value] = v.Index
	}
}

// ActionText gets text of damage message by its index.
func ActionText(lang Lang, index action.Damage) string {
	return translationByIndex[lang][index]
}

// ActionIndex gets index of damage message by its text for provided
// language.
func ActionIndex(lang Lang, text string) action.Damage {
	return translationByText[lang][text]
}
