package main

import (
	"fmt"

	"github.com/valeri2000/go-simple-scheduler/job"
)

func printHiThere() {
	fmt.Println("hi there")
}

func main() {
	fmt.Println("main()")
	testJob := job.CreateJob(true, 0, printHiThere)
	testJob.Run()
}
