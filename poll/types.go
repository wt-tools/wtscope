package poll

import "net/http"

type httper interface {
	Do(req *http.Request) (*http.Response, error)
}

type task struct {
	method, url string
	repeat      int
	retry       int
	ret         chan []byte
}
