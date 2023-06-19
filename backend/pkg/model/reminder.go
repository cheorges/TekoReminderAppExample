package model

type Reminder struct {
	Id   int
	Task string
	Done bool
}

func New(task string) Reminder {
	return Reminder{Task: task}
}

func (r *Reminder) Toggle() {
	r.Done = !r.Done
}
