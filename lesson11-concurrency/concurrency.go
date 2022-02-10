package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	goroutines = 10
)

func FillArray(array []int) {
	fillPart(array, len(array), 0)
}

func ConcurrentFillArray(array []int) {
	partLen := len(array) / goroutines
	var wg sync.WaitGroup

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(v int) {
			fillPart(array, partLen, v)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func fillPart(array []int, partLen, v int) {
	var start, end int

	start = partLen * v
	if v == goroutines-1 {
		end = len(array)
	} else {
		end = partLen * (v + 1)
	}

	for i := start; i < end; i++ {
		array[i] = i
	}
}

func main() {
	arr := make([]int, 1000000)

	cft := time.Now()
	ConcurrentFillArray(arr)
	fmt.Printf("Concurrent: %v\n", time.Since(cft))

	arr2 := make([]int, 1000000)

	ft := time.Now()
	FillArray(arr2)
	fmt.Printf("Non-concurrent: %v\n", time.Since(ft))
}
