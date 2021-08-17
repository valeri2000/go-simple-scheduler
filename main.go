package main

import (
	"fmt"
	"time"

	"github.com/valeri2000/go-simple-scheduler/job"
	"github.com/valeri2000/go-simple-scheduler/scheduler"
)

func printHiThere() {
	fmt.Println("hi there |", time.Now())
}

func fib(n int) {
	a := 1
	b := 1
	for i := 3; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	fmt.Println("fib ", n, "=", b)
}

func add(a int, b int) {
	fmt.Printf("add(%d, %d) = %d | %s\n",
		a,
		b,
		a+b,
		time.Now())
}

func main() {
	add2plus3 := func() {
		add(2, 3)
	}
	add5plus5 := func() {
		add(5, 5)
	}
	fib5 := func() {
		fib(5)
	}

	scheduler := scheduler.CreateScheduler()
	scheduler.AddJob(job.CreateJob(false, 3, printHiThere))
	scheduler.AddJob(job.CreateJob(false, 1, add2plus3))
	scheduler.AddJob(job.CreateJob(true, 0, fib5)) // run once
	scheduler.AddJob(job.CreateJob(false, 2, add5plus5))
	scheduler.Start()
}
