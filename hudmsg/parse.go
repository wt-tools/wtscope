package hudmsg

import (
	"fmt"

	"github.com/wt-tools/adjutant/action"
)

type (
	token struct {
		pos   int
		index tokenType
		text  []rune
	}
	tokenType int
)

const (
	unknownTok tokenType = iota
	clanTagTok
	playerNameTok
	vehicleTok
	actionTok
	achievementTok
)

// Just funny to parse it without regexps.
func parseDamage(msg string, id uint) (action.Damage, error) {
	// The message started with username possibly prepended by clan tag.
	var (
		mode           tokenType = playerNameTok
		curToken       []rune
		quotes, parens counter
		tokens         []token
	)
	for i, c := range msg {
		curToken = append(curToken, c)
		if parens.insideParens(c) {
			continue
		}
		if quotes.insideQuotes(c) {
			continue
		}
		if !isSpace(c) {
			continue
		}
		newToken := token{pos: i, text: curToken}
		fmt.Printf("raw token: %s\n", string(curToken))
		switch mode {
		case clanTagTok, playerNameTok:
			if parens.occurs {
				mode = vehicleTok
			}
		case vehicleTok:
			if !parens.occurs {
				mode = actionTok
			}
		case actionTok:
			mode = playerNameTok
		case achievementTok:
			if quotes.occurs {
				mode = achievementTok
			}
		}
		quotes.reset()
		parens.reset()
		curToken = nil
		newToken.index = mode
		tokens = append(tokens, newToken)
	}
	for _, m := range tokens {
		fmt.Printf("%d: %d %v\n", m.pos, m.index, string(m.text))
	}
	return action.Damage{Origin: msg, ID: id}, nil
}

func isSpace(c rune) bool {
	if c == ' ' {
		return true
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
