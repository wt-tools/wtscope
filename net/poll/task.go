package poll

import "os"

type Task struct {
	name, method, url string
	repeat            int
	retry             int
	log               *os.File
	ret               chan []byte
}

// Update URL for task.
func (t *Task) Update(url string) {
	t.url = url
}

// Returns channel that provides task results.
func (t *Task) Results() chan []byte {
	return t.ret
}
