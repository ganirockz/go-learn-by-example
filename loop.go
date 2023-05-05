package main

import "fmt"

// for is Go’s only looping construct
func main() {

	// The most basic type, with a single condition.
	fmt.Println("The most basic type, with a single condition.")
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// A classic initial/condition/after for loop.
	fmt.Println("A classic initial/condition/after for loop.")
	for j := 1; j < 4; j++ {
		fmt.Println(j)
	}

	// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	fmt.Println("for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.")
	for {
		fmt.Println("Loop")
		break
	}

	// You can also continue to the next iteration of the loop.
	for j := 2; j < 5; j++ {
		if j == 4 {
			continue
		}

		fmt.Println(j)
	}

}
