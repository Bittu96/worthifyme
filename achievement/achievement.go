package achievement

import "time"

type Achievement struct {
	task     string
	status   bool
	created  time.Time
	finished time.Time
}

func (a *Achievement) Set() Achievement {
	a.finished = time.Now()
}

func New(task string) Achievement {
	return Achievement{
		task:    task,
		status:  false,
		created: time.Now(),
	}
}
