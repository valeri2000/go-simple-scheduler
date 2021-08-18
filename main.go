package main

import (
	"fmt"
	"time"

	"github.com/valeri2000/go-simple-scheduler/job"
	"github.com/valeri2000/go-simple-scheduler/scheduler"
)

func printHiThere() {
	fmt.Println("hi there |", time.Now().UnixNano()/int64(time.Millisecond))
}

func fib(n int) {
	a := 1
	b := 1
	for i := 3; i <= n; i++ {
		c := a + b
		a = b
		b = c
	}
	fmt.Println("fib ", n, "=", b, "|", time.Now().UnixNano()/int64(time.Millisecond))
}

func add(a int, b int) {
	fmt.Printf("add(%d, %d) = %d | %d\n",
		a,
		b,
		a+b,
		time.Now().UnixNano()/int64(time.Millisecond))
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

	scheduler, err := scheduler.CreateScheduler("millisecond")
	//scheduler, err := scheduler.CreateScheduler("second")
	if err != nil {
		panic("Error")
	}

	scheduler.AddJob(job.CreateJob(3, 300, printHiThere))
	scheduler.AddJob(job.CreateJob(0, 100, add2plus3))
	scheduler.AddJob(job.CreateJob(1, 0, fib5)) // run once
	scheduler.AddJob(job.CreateJob(0, 200, add5plus5))
	scheduler.Start()
}
