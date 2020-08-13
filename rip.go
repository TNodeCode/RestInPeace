package main

import (
	"flag"
	"fmt"
	"math"
	"sync"
)

type Result struct {
	n     int
	prime bool
}

func main() {
	var nMax = flag.Int("n", 10000, "Maximum number")
	var nThreads = flag.Int("t", 1, "Number of threads")
	flag.Parse()

	fmt.Println("Number of Threads: ", *nThreads)
	fmt.Println("Maximum Number: ", *nMax)

	testNumbersParallel(*nMax, *nThreads)
}

func testNumbersParallel(nMax int, threads int) {
	wg := &sync.WaitGroup{}
	flag.Parse()

	c := make(chan Result)
	countPrimes := 0
	nCurrent := 2

	// Start threads
	for x := 2; x < 2+threads+1; x++ {
		wg.Add(1)

		// Start new Go routine
		go isPrimeAsync(x, c, wg)

		nCurrent++
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for res := range c { // Wait for the channel to return some value
		if res.prime {
			countPrimes++ // Increment prime counter
		}

		if nCurrent < nMax { // If limit is not reached start new goroutine
			wg.Add(1)
			go isPrimeAsync(nCurrent, c, wg)
			nCurrent++
		}
	}

	fmt.Println("---\n", countPrimes, "prime numbers found")
}

func isPrimeAsync(i int, c chan Result, wg *sync.WaitGroup) {
	// Primitive algorithm for calculating prime numbers
	for j := 2; j < int(math.Sqrt(float64(i)))+1; j++ {
		if i%j == 0 { // not a prime
			// Post result to the channel
			c <- Result{i, false}
			defer wg.Done()
			return
		}
	}

	// Post result to the channel
	c <- Result{i, true}
	defer wg.Done()
}
