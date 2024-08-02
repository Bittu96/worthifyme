package progress

import "worthifyme/achievement"

type Progress struct {
	status              string
	totalAchievements   int
	currentAchievements int
	achievements        []interface{}
}

func (p *Progress) AddAchievement(a achievement.Achievement) {
	p.achievements = append(p.achievements, a)
}

func (p *Progress) Start() {
	p.status = "inProgress"
}

func New(t int) Progress {
	return Progress{
		totalAchievements:   t,
		currentAchievements: 0,
		status:              "created",
	}
}
