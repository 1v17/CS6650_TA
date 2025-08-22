package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	goroutines = 50
	iterations = 1000
)

func plainMapTest() {
	fmt.Println("Plain map test:")
	m := make(map[int]int)
	var wg sync.WaitGroup
	start := time.Now()
	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				m[g*iterations+i] = i // Not safe!
			}
		}(g)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("len(m) = %d, time = %v\n", len(m), elapsed)
	// Likely to crash due to concurrent map writes!
}

type MutexMap struct {
	m  map[int]int
	mu sync.Mutex
}

func mutexMapTest() {
	fmt.Println("Mutex map test:")
	mm := &MutexMap{m: make(map[int]int)}
	var wg sync.WaitGroup
	start := time.Now()
	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				mm.mu.Lock()
				mm.m[g*iterations+i] = i
				mm.mu.Unlock()
			}
		}(g)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("len(m) = %d, time = %v\n", len(mm.m), elapsed)
}

type RWMutexMap struct {
	m  map[int]int
	mu sync.RWMutex
}

func rwMutexMapTest() {
	fmt.Println("RWMutex map test:")
	rm := &RWMutexMap{m: make(map[int]int)}
	var wg sync.WaitGroup
	start := time.Now()
	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				rm.mu.Lock()
				rm.m[g*iterations+i] = i
				rm.mu.Unlock()
			}
		}(g)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("len(m) = %d, time = %v\n", len(rm.m), elapsed)
}

func syncMapTest() {
	fmt.Println("sync.Map test:")
	var m sync.Map
	var wg sync.WaitGroup
	start := time.Now()
	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				m.Store(g*iterations+i, i)
			}
		}(g)
	}
	wg.Wait()
	count := 0
	m.Range(func(_, _ any) bool {
		count++
		return true
	})
	elapsed := time.Since(start)
	fmt.Printf("len(m) = %d, time = %v\n", count, elapsed)
}

func Collection() {
	// Run each test 3 times and average results manually
	// plainMapTest() // Expect crash!
	mutexMapTest()
	rwMutexMapTest()
	syncMapTest()
}
