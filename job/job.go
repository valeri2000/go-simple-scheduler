package job

type Job struct {
	RunOnce    bool
	Period     int // in seconds
	TimeLeft   int
	PendingRun bool
	Func       func()
}

func CreateJob(RunOnce bool, Period int, Func func()) *Job {
	return &Job{RunOnce: RunOnce,
		Period:     Period,
		TimeLeft:   Period,
		PendingRun: (Period == 0),
		Func:       Func,
	}
}

func (job *Job) ShouldRun() bool {
	return job.PendingRun
}

func (job *Job) DecreaseTime() {
	job.TimeLeft -= 1
}

func (job *Job) ResetTime() bool {
	if job.RunOnce == true {
		return false
	}

	job.TimeLeft = job.Period
	job.PendingRun = (job.Period == 0)
	return true
}

func (job *Job) Run() {
	go job.Func()
	go job.ResetTime()
}
