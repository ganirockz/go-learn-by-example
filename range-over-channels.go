package main

import "fmt"

// In a previous examples we saw how for and range provide iteration over basic data structures. We can also use this syntax to iterate over values received from a channel.
func main() {

	// We’ll iterate over 2 values in the queue channel.
	queue := make(chan int, 2)
	queue <- 1
	queue <- 2
	close(queue)

	// This range iterates over each element as it’s received from queue. Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}
