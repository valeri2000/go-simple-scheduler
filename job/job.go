package job

type Job struct {
	RunOnce bool
	Period  int // in seconds
	Func    func()
}

func CreateJob(RunOnce bool, Period int, Func func()) *Job {
	return &Job{RunOnce: RunOnce, Period: Period, Func: Func}
}

func (job *Job) Run() {
	job.Func()
}
