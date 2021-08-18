package scheduler

import (
	"fmt"
	"time"

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

func (scheduler *Scheduler) Tick() {
	time.Sleep(1 * time.Second)
	for i := range scheduler.Jobs {
		scheduler.Jobs[i].DecreaseTime()
	}
}

func (scheduler *Scheduler) RunPending() {
	for i := range scheduler.Jobs {
		if scheduler.Jobs[i].ShouldRun() {
			go scheduler.Jobs[i].Run()
			go scheduler.Jobs[i].ResetTime()
		}
	}
}

func (scheduler *Scheduler) Start() {
	fmt.Println("Start() |", time.Now())
	func() {
		for {
			go scheduler.RunPending()
			scheduler.Tick()
		}
	}()
}
