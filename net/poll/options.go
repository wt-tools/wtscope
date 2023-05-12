package poll

import (
	"net/http"
	"time"
)

type option func(*Service)

func SetLoopDelay(d time.Duration) option {
	return func(s *Service) {
		s.loopDelay = d
	}
}

func SetProblemDelay(d time.Duration) option {
	return func(s *Service) {
		s.onProblemDelay = d
	}
}

func SetClient(c *http.Client) option {
	return func(s *Service) {
		s.httpc = c
	}
}

func SetLogger(l chan error) option {
	return func(s *Service) {
		s.err = l
	}
}
