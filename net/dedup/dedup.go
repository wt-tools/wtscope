package dedup

import (
	"bytes"
	"time"
)

type (
	duplicate struct {
		data         []byte
		prev, latest uint
		prevAt       time.Time
	}
	item struct {
		val uint
		at  time.Time
	}
)

func New() *duplicate {
	return &duplicate{}
}

// Block duplicated lines of content.
func (s *duplicate) BlockContent(data []byte) bool {
	if bytes.Equal(data, s.data) {
		return true
	}
	s.data = data
	return false
}

func (s *duplicate) Exists(val uint) bool {
	const shieldPeriod = 30 * time.Second
	// new val should be greater than keeped one
	if val <= s.latest {
		return true
	}
	// after shield period shift values
	if time.Since(s.prevAt) > shieldPeriod {
		s.prev = s.latest
		s.prevAt = time.Now()
	}
	s.latest = val
	// queue probably restarted
	if s.latest < s.prev {
		s.latest = 0
		s.prev = 0
		s.prevAt = time.Now()
	}
	return false
}
