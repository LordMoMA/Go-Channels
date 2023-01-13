package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest(r chan<- int32) {
	// Simulate a workload.
	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano()) // before Go 1.20

	// ra, rb := make(chan int32), make(chan int32)
	// go longTimeRequest(ra)
	// go longTimeRequest(rb)

	// fmt.Println(sumSquares(<-ra, <-rb))

	// The channel can be buffered or not.
	results := make(chan int32, 2)
	go longTimeRequest(results)
	go longTimeRequest(results)

	fmt.Println(sumSquares(<-results, <-results))
}
