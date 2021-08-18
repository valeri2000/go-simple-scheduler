package job

import (
	"time"
)

type Job struct {
	TimesToRun  uint  // 0 -> no limit
	Period      int64 // in ns
	LastRunTime int64
	Func        func()
}

func CreateJob(TimesToRun uint, Period time.Duration, Func func()) *Job {
	return &Job{
		TimesToRun:  TimesToRun,
		Period:      int64(Period),
		LastRunTime: time.Now().UnixNano(),
		Func:        Func,
	}
}

func (job *Job) Run() {
	for {
		if time.Now().UnixNano()-job.LastRunTime >= job.Period {
			go job.Func()
			job.LastRunTime = time.Now().UnixNano()

			if job.TimesToRun == 1 {
				break
			}
			if job.TimesToRun > 1 {
				job.TimesToRun -= 1
			}
		}
	}
}
