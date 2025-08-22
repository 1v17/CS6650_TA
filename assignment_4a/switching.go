package main

import (
	"fmt"
	"runtime"
	"time"
)

const rounds = 1_000_000

func measureSwitchTime(setProcs int) time.Duration {
	runtime.GOMAXPROCS(setProcs)
	ch := make(chan struct{})
	done := make(chan struct{})

	start := time.Now()

	go func() {
		for i := 0; i < rounds; i++ {
			ch <- struct{}{}
			<-ch
		}
		done <- struct{}{}
	}()

	go func() {
		for i := 0; i < rounds; i++ {
			<-ch
			ch <- struct{}{}
		}
	}()

	<-done
	elapsed := time.Since(start)
	return elapsed
}

func SwitchingMain() {
	// Single OS thread
	single := measureSwitchTime(1)
	fmt.Printf("GOMAXPROCS=1:   total=%v, avg switch=%v\n", single, single/(2*rounds))

	// Multiple OS threads (default)
	multi := measureSwitchTime(runtime.NumCPU())
	fmt.Printf("GOMAXPROCS=%d: total=%v, avg switch=%v\n", runtime.NumCPU(), multi, multi/(2*rounds))
}
