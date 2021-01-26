package main

import (
	"flag"
	"fmt"
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

/**
 * Function is responsible for managing parallel Go routines
 */
func testNumbersParallel(nMax int, nParallel int) {
	wg := &sync.WaitGroup{}
	flag.Parse()

	c := make(chan Result)
	countPrimes := 0 // Number of primes that have been found
	nCurrent := 2    // Currently checked number

	// Start nParallel Go routines
	for x := 2; x < 2+nParallel+1; x++ {
		// Increment WaitGroup counter
		wg.Add(1)

		// Start new Go routine
		go isPrime(x, c, wg)

		// Increment current number that is checked for being prime
		nCurrent++
	}

	go func() {
		// Block until the WaitGroup counter goes back to 0 and then close the channel.
		// This will happen only once, because every time a Go routine finishes and the counter
		// decrements a new Go routine will be started and the counter gets incremented again
		// as long as not all number have been checked
		wg.Wait()
		close(c)
	}()

	// Wait for the channel to return some value
	// This code blocks until the channel is empty
	for res := range c {
		if res.prime {
			countPrimes++ // Increment prime counter
		}

		// If limit is not reached start new goroutine
		if nCurrent < nMax {
			wg.Add(1)                   // Increment counter ofWaitGroup
			go isPrime(nCurrent, c, wg) // Run calculation in Go routine
			nCurrent++
		}
	}

	// This code will run when the channel is empty
	fmt.Println("---\n", countPrimes, "prime numbers found")
}

/**
 * Function that checks if a given number is a prime number
 */
func isPrime(i int, c chan Result, wg *sync.WaitGroup) {
	// Simple algorithm for calculating prime numbers
	for j := 2; j < i; j++ {
		if i%j == 0 { // not a prime
			// Post result to the channel
			c <- Result{i, false}

			// Decrement WaitGroup counter as soon as the function execution is done
			defer wg.Done()
			return
		}
	}

	// Post result to the channel
	c <- Result{i, true}

	// Decrement WaitGroup counter as soon as the function execution is done
	defer wg.Done()
}
