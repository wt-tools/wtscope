package hudmsg

//go:generate stringer -type=tokenType

import (
	"strings"
	"time"

	"github.com/wt-tools/wtscope/action"
	"github.com/wt-tools/wtscope/events"
	"github.com/wt-tools/wtscope/l10n"
	"github.com/wt-tools/wtscope/vehicle"
)

type (
	token struct {
		pos   int
		index tokenType
		text  string
	}
	tokenType int
)

const weirdDisconnectMessage = `kd?NET_PLAYER_DISCONNECT_FROM_GAME`
const (
	unknownType tokenType = iota
	squadTagType
	playerNameType
	vehicleType
	actionType
	achievementType
)

// Just funny to parse it without regexps.
func parseDamage(dmg Damage) (events.Event, error) {
	tokens := tokenize(dmg)
	var (
		p1, p2                 events.Player
		v1, v2                 vehicle.Vehicle
		rawAct                 strings.Builder
		achiev                 string
		actionCode, lastAction action.Code
		lang                   l10n.Lang
		actionComplete         bool
	)
	for _, tok := range tokens {
		//	loop:
		switch tok.index {
		case squadTagType:
			if p1.Name == "" {
				p1.Squad = string(tok.text)
				break
			}
			if p2.Name == "" {
				p2.Squad = string(tok.text)
			}
		case playerNameType:
			if p1.Name == "" {
				p1.Name = string(tok.text)
				break
			}
			if p2.Name == "" {
				p2.Name = string(tok.text)
				break
			}
		case vehicleType:
			if v1.Name == "" {
				v1.Name = string(tok.text)
				break
			}
			if v2.Name == "" {
				v2.Name = string(tok.text)
				break
			}
		case actionType:
			// if actionComplete {
			//	// Convert actions to another kinds of types.
			//	switch lastAction {
			//	case action.Destroyed:
			//		tok.index = playerNameType
			//	}
			//	goto loop
			// }
			if rawAct.Len() > 0 {
				rawAct.WriteString(" ")
			}
			rawAct.WriteString(tok.text)
			// If known action has found. Treat all
			// rest of actions tokens as player/vehicle/achievement.
			actionCode, lang, actionComplete = l10n.FindAction(rawAct.String())
			if actionComplete {
				lastAction = actionCode
			}
		case achievementType:
			achiev = string(tok.text)
		}
	}
	d := events.Event{
		ID:            dmg.ID,
		Origin:        dmg.Msg,
		Lang:          lang,
		Enemy:         dmg.Enemy,
		Player:        p1,
		Vehicle:       v1,
		TargetPlayer:  p2,
		TargetVehicle: v2,
		Action:        lastAction,
		ActionText:    rawAct.String(),
		Achievement: &events.Achievement{
			Name: achiev,
		},
		At: time.Duration(dmg.Time) * time.Second,
	}
	// Patch result if the weird message occured.
	if d.ActionText == weirdDisconnectMessage {
		d.Action = action.NetworkDisconnect
		d.ActionText = l10n.ActionText(l10n.En, action.Disconnected)
		if len(d.Player.Name) > 3 {
			d.Player.Name = d.Player.Name[:len(d.Player.Name)-3]
		}
	}
	return d, nil
}

func tokenize(dmg Damage) []token {
	// The message started with username possibly prepended by clan tag.
	var (
		mode           = playerNameType
		word           []rune
		quotes, parens counter
		prevToken      = &token{}
		tokens         []token
	)
	for i, c := range dmg.Msg + " " { // TODO cleanup hack with trailing space
		if parens.insideParens(c) || quotes.insideQuotes(c) {
			word = append(word, c)
			continue
		}
		if !isSpace(c) {
			word = append(word, c)
			continue
		}
		newToken := token{pos: i, text: string(word)}
		switch mode {
		case squadTagType:
			if !isSquadTag(word) {
				mode = playerNameType
			}
		case playerNameType:
			if parens.occurs {
				mode = vehicleType
				break
			}
			if isSquadTag(word) {
				mode = squadTagType
				break
			}
			if prevToken.index == playerNameType {
				mode = actionType
			}
		case vehicleType:
			if !parens.occurs {
				mode = actionType
			}
		case actionType:
			if quotes.occurs {
				mode = achievementType
				break
			}
			if parens.occurs {
				// second player + vehicle info
				mode = vehicleType
				prevToken.index = playerNameType
				break
			}
			if isSquadTag(word) {
				mode = squadTagType
			}
		}
		// cleanup quotes and parens
		if mode == vehicleType || mode == achievementType {
			newToken.text = string(word[1 : len(word)-1])
		}
		newToken.index = mode
		tokens = append(tokens, newToken)
		prevToken = &tokens[len(tokens)-1]
		quotes.reset()
		parens.reset()
		word = nil
	}
	return tokens
}

func isSpace(c rune) bool {
	if c == ' ' {
		return true
	}
	return false
}

func isSquadTag(s []rune) bool {
	if len(s) < 2 {
		return false
	}
	// squad prefixes
	prefixes := []rune{'[', '^', '-', '=', '⋇', '.'}
	suffixes := []rune{']', '^', '-', '=', '⋇', '.'}
	for i, t := range prefixes {
		if s[0] == t && s[len(s)-1] == suffixes[i] {
			return true
		}
	}
	return false
}

type counter struct {
	val    int
	occurs bool
}

func (c *counter) reset() {
	c.val = 0
	c.occurs = false
}

func (c *counter) insideQuotes(r rune) bool {
	if r == '"' {
		c.occurs = true
		c.val++
	}
	if c.val == 0 || c.val >= 2 {
		return false
	}
	return true
}

func (c *counter) insideParens(r rune) bool {
	if r == '(' {
		c.occurs = true
		c.val++
		return true
	}
	if r == ')' {
		c.val--
		// it takes into account case for unbalanced parensis
		if c.val <= 0 {
			return false
		}
		return true
	}
	if c.val > 0 {
		return true
	}
	return false
}
