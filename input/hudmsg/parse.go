package hudmsg

//go:generate stringer -type=tokenType

import (
	"fmt"
	"time"

	"github.com/wt-tools/wtscope/action"
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

const (
	unknownType tokenType = iota
	clanTagType
	playerNameType
	vehicleType
	actionType
	achievementType
)

// Just funny to parse it without regexps.
func parseDamage(dmg Damage) (action.GeneralAction, error) {
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
		case clanTagType:
			if !isClanTag(word) {
				mode = playerNameType
			}
		case playerNameType:
			if parens.occurs {
				mode = vehicleType
				break
			}
			if isClanTag(word) {
				mode = clanTagType
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
			// second player + vehicle info
			if parens.occurs {
				mode = vehicleType
				prevToken.index = playerNameType
				break
			}
			if isClanTag(word) {
				mode = clanTagType
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
	for _, m := range tokens {
		fmt.Printf("%d: %s :: %v\n", m.pos, m.index.String(), string(m.text))
	}
	var (
		p1, p2 action.Player
		v1, v2 vehicle.Vehicle
		act    action.Action
	)
	for _, tok := range tokens {
		switch tok.index {
		case clanTagType:
			if p1.Clan == "" {
				p1.Clan = string(tok.text)
				break
			}
			if p2.Clan == "" {
				p2.Clan = string(tok.text)
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
			// XXX
		}
	}

	d := action.GeneralAction{
		ID:     dmg.ID,
		Origin: dmg.Msg,
		Enemy:  dmg.Enemy,
		Damage: &action.Damage{
			Player:        p1,
			Vehicle:       v1,
			TargetPlayer:  p2,
			TargetVehicle: v2,
			Action:        act,
		},
		At: time.Duration(dmg.Time) * time.Second,
	}
	return d, nil
}

func isSpace(c rune) bool {
	if c == ' ' {
		return true
	}
	return false
}

func isClanTag(s []rune) bool {
	if len(s) < 2 {
		return false
	}
	prefixes := []rune{'[', '^', '-', '=', '⋇'}
	suffixes := []rune{']', '^', '-', '=', '⋇'}
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
