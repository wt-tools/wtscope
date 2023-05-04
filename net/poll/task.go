package poll

type Task struct {
	method, url string
	repeat      int
	retry       int
	ret         chan []byte
}

// Update URL for task.
func (t *Task) Update(url string) {
	t.url = url
}

// Returns channel that provides task results.
func (t *Task) Results() chan []byte {
	return t.ret
}
