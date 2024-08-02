package reward

import (
	"time"
	"worthifyme/progress"
)

type reward struct {
	title       string
	description string
	status      string
	progress    progress.Progress
	created     time.Time
}

func (r reward) Get() *reward {
	return &r
}

func (r *reward) AddProgress(n int) {
	r.progress = progress.New(n)
}

func (r *reward) StartProgress() {
	r.status = "inProgress"
	r.progress.Start()
}

func (r *reward) DeleteProgress() {
	r.status = "stopped"
	r.progress = progress.Progress{}
}

func New(title, description string) reward {
	return reward{
		title:       title,
		description: description,
		status:      "created",
		created:     time.Now(),
	}
}
