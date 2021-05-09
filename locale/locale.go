// Package locale defines index for tokens alongside with their
// translations.
package locale

import "github.com/wt-tools/adjutant/damage"

type Translation struct {
	Lang  Lang
	Index damage.Action
	Value string
}

var (
	translationByIndex = make(map[Lang]map[damage.Action]string)
	translationByText  = make(map[Lang]map[string]damage.Action)
)

func init() {
	for _, v := range actionTexts {
		if _, ok := translationByIndex[v.Lang]; !ok {
			translationByIndex[v.Lang] = make(map[damage.Action]string)
		}
		translationByIndex[v.Lang][v.Index] = v.Value
		if _, ok := translationByText[v.Lang]; !ok {
			translationByIndex[v.Lang] = make(map[damage.Action]string)
		}
		translationByText[v.Lang][v.Value] = v.Index
	}
}

// ActionText gets text of damage message by its index.
func ActionText(lang Lang, index damage.Action) string {
	return translationByIndex[lang][index]
}

// ActionIndex gets index of damage message by its text for provided
// language.
func ActionIndex(lang Lang, text string) damage.Action {
	return translationByText[lang][text]
}
