// Package l10n defines translations of words from WT logs.
package l10n

import "github.com/wt-tools/wtscope/action"

type translation struct {
	Lang  Lang
	Value string
}
type codeLang struct {
	Code action.Code
	Lang Lang
}

var (
	actionsByCode     = make(map[Lang]map[action.Code]string)
	actionsByText     = make(map[Lang]map[string]action.Code)
	actionsByTextOnly = make(map[string]codeLang)
)

func init() {
	for _, v := range actionTexts {
		for _, t := range v.List {
			if _, ok := actionsByCode[t.Lang]; !ok {
				actionsByCode[t.Lang] = make(map[action.Code]string)
			}
			actionsByCode[t.Lang][v.Code] = t.Value
			if _, ok := actionsByText[t.Lang]; !ok {
				actionsByText[t.Lang] = make(map[string]action.Code)
			}
			actionsByText[t.Lang][t.Value] = v.Code
			actionsByTextOnly[t.Value] = codeLang{v.Code, t.Lang}
		}
	}
}

// ActionText gets text of damage message by its index.
func ActionText(lang Lang, code action.Code) string {
	return actionsByCode[lang][code]
}

// ActionIndex gets index of damage message by its text for provided
// language.
func ActionIndex(lang Lang, text string) action.Code {
	return actionsByText[lang][text]
}

// FindAction finds action in any language.
func FindAction(text string) (action.Code, Lang, bool) {
	v, ok := actionsByTextOnly[text]
	return v.Code, v.Lang, ok
}
