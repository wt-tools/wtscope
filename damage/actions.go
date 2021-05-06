package damage

import "github.com/wt-tools/adjutant/locale"

type Action struct {
	Name string
	ID   int
}

const (
	Unknown = iota
	Destroyed
	Wrecked
	Fired
)

var actionRegistry = map[locale.Lang][]string{
	locale.En: []string{"unknown action", "destroyed", "has been wrecked", "?"},
	locale.Ru: []string{"неизвестное действие", "уничтожил", "выведен из строя", "поджег"},
}

func Parse(lang locale.Lang, action string) (Action, error) {
	return Action{}, errNotImplemented
}

// name ...
func ActionByName(lang locale.Lang, name string) int {
	return 0
}

// ActionID ...
func ActionByID(lang locale.Lang, id int) (string, error) {
	if len(actionRegistry[lang]) > id+1 {
		return "", errUnknownActionID
	}
	return actionRegistry[lang][id], nil
}
