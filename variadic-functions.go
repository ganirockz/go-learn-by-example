package main

import "fmt"

// Here’s a function that will take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Println("nums: ", nums)

	total := 0

	// Within the function, the type of nums is equivalent to []int. We can call len(nums), iterate over it with range, etc.
	for _, num := range nums {
		total += num
	}

	fmt.Println("Total: ", total)
}

// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.
func main() {

	sum(1, 2)

	sum(1, 2, 3)

	// If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.
	nums := []int{1, 2, 3, 4}

	sum(nums...)
}
