package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel c
}

func main() {
	fmt.Println(">> DAY 9: CONCURRENCY & CHANNELS")

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	// Start two goroutines to split the work
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// Receive the results from the channel
	x, y := <-c, <-c

	fmt.Println("Result X, Y, and Total:", x, y, x+y)
}

