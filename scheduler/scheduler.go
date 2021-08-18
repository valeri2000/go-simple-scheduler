package scheduler

import (
	"errors"
	"fmt"
	"time"

	"github.com/valeri2000/go-simple-scheduler/job"
)

type Scheduler struct {
	Precision time.Duration // time unit
	Jobs      []*job.Job
}

func CreateScheduler(precisionStr string) (*Scheduler, error) {
	switch precisionStr {
	case "second":
		return &Scheduler{Precision: time.Second, Jobs: make([]*job.Job, 0)}, nil
	case "millisecond":
		return &Scheduler{Precision: time.Millisecond, Jobs: make([]*job.Job, 0)}, nil
	default:
		return nil, errors.New("Precision not supported!")
	}
}

func (scheduler *Scheduler) AddJob(newJob *job.Job) {
	scheduler.Jobs = append(scheduler.Jobs, newJob)
}

func (scheduler *Scheduler) Tick() {
	time.Sleep(1 * scheduler.Precision)
	for i := range scheduler.Jobs {
		scheduler.Jobs[i].DecreaseTime()
	}
}

func (scheduler *Scheduler) RunPending() {
	for i := range scheduler.Jobs {
		if scheduler.Jobs[i].ShouldRun() { // TODO handle done jobs better
			go scheduler.Jobs[i].Run()
			go scheduler.Jobs[i].ResetTime()
		}
	}
}

func (scheduler *Scheduler) Start() {
	fmt.Println("Start() |", time.Now())
	for {
		go scheduler.RunPending()
		scheduler.Tick()
	}
}
