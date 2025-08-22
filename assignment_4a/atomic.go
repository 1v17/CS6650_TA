package main

import (
	"fmt"
	"sync"
	// "sync/atomic"
)

func AtomicMain() {

	// var ops atomic.Uint64
	var ops int64

	var wg sync.WaitGroup

	for range 50 {
		wg.Add(1)

		go func() {
			for range 1000 {

				// ops.Add(1)
				ops++
			}

			wg.Done()
		}()
	}

	wg.Wait()

	// fmt.Println("ops:", ops.Load())
	fmt.Println("ops:", ops)
}
