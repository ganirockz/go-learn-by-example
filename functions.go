package main

import "fmt"

// Here’s a function that takes two ints and returns their sum as an int.
func plus(a int, b int) int {
	return a + b
}

// When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Call a function just as you’d expect, with name(args).
func main() {
	a := 1
	b := 2

	c := plus(a, b)

	fmt.Println("Sum of ", a, " and", b, " is ", c)

	d := plusPlus(a, b, c)

	fmt.Println("D is: ", d)

}
