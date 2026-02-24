package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	value int
	N     int
}

func writeData(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

func sum(numChan chan int, resChan chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range numChan {
		sum := 0
		for j := 1; j <= v; j++ {
			sum += j
		}
		resChan <- Result{sum, v}
	}
}

func main() {
	numChan := make(chan int, 2000)
	resChan := make(chan Result, 2000)
	results := make([]int, 2001)

	go writeData(numChan)

	var wg sync.WaitGroup

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go sum(numChan, resChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	for v := range resChan {
		results[v.N] = v.value
	}

	for i := 1; i <= 2000; i++ {
		fmt.Printf("results[%d]=%d\n", i, results[i])
		time.Sleep(10 * time.Millisecond)
	}
}
