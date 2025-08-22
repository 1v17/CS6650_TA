package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	filename       = "output.txt"
	fileIterations = 100000
	lineToWrite    = "This is a line of text.\n"
	bufferedFile   = "output_buffered.txt"
	unbufferedFile = "output_unbuffered.txt"
)

func writeUnbuffered() time.Duration {
	f, err := os.Create(unbufferedFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	start := time.Now()
	for i := 0; i < fileIterations; i++ {
		_, err := f.Write([]byte(lineToWrite))
		if err != nil {
			panic(err)
		}
	}
	elapsed := time.Since(start)
	return elapsed
}

func writeBuffered() time.Duration {
	f, err := os.Create(bufferedFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	start := time.Now()
	for i := 0; i < fileIterations; i++ {
		_, err := w.WriteString(lineToWrite)
		if err != nil {
			panic(err)
		}
	}
	w.Flush()
	elapsed := time.Since(start)
	return elapsed
}

func FileMain() {
	unbuf := writeUnbuffered()
	fmt.Printf("Unbuffered write: %v\n", unbuf)

	buf := writeBuffered()
	fmt.Printf("Buffered write:   %v\n", buf)
}
