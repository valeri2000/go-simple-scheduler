package scheduler

import (
	"github.com/valeri2000/go-simple-scheduler/job"
)

type Scheduler struct {
	Jobs []*job.Job
}

func CreateScheduler() *Scheduler {
	return &Scheduler{Jobs: make([]*job.Job, 0)}
}

func (scheduler *Scheduler) AddJob(newJob *job.Job) {
	scheduler.Jobs = append(scheduler.Jobs, newJob)
}

func (scheduler *Scheduler) Start() {
	for _, job := range scheduler.Jobs {
		go job.Run()
	}
	for {
	}
}
