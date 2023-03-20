package main

import (
	"fmt"
)

// sum the numbers in a and send the result on res.
func sum(a []int, res chan<- int) {
	sum := 0
	for _, number := range a {
		sum += number
	}
	res <- sum
}

// concurrently sum the array a.
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)

	// TODO Get the subtotals from the channel and return their sum
	sum := <-ch
	sum += <-ch

	return sum
}

func main() {
	// example call
	a := []int{1, 2, 3, 4}
	fmt.Println(ConcurrentSum(a))
}
