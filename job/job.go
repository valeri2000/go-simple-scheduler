package job

type Job struct {
	TimesToRun uint // 0 -> no limit
	Done       bool
	Period     uint // in seconds
	TimeLeft   uint
	PendingRun bool
	Func       func()
}

func CreateJob(TimesToRun uint, Period uint, Func func()) *Job {
	return &Job{
		TimesToRun: TimesToRun,
		Done:       false,
		Period:     Period,
		TimeLeft:   Period,
		PendingRun: (Period == 0),
		Func:       Func,
	}
}

func (job *Job) ShouldRun() bool {
	return !job.Done && job.PendingRun
}

func (job *Job) DecreaseTime() {
	if !job.Done {
		job.TimeLeft -= 1
		if job.TimeLeft < 1 {
			job.PendingRun = true
		}
	}
}

func (job *Job) ResetTime() {
	if job.TimesToRun == 0 {
		job.TimeLeft = job.Period
		job.PendingRun = (job.Period == 0)
	} else if job.TimesToRun == 1 {
		job.Done = true
	} else {
		job.TimesToRun -= 1
		job.TimeLeft = job.Period
		job.PendingRun = (job.Period == 0)
	}
}

func (job *Job) Run() {
	go job.Func()
}
