package main

import "fmt"

// We’ll show how pointers work in contrast to values with 2 functions: zeroval and zeroptr. zeroval has an int parameter, so arguments will be passed to it by value. zeroval will get a copy of ival distinct from the one in the calling function.
func zeroVal(val int) {
	val = 0
}

// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer. The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address. Assigning a value to a dereferenced pointer changes the value at the referenced address.
func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1

	fmt.Println("Intial value: ", i)

	zeroVal(i)
	fmt.Println(i)

	// zeroval doesn’t change the i in main, but zeroptr does because it has a reference to the memory address for that variable.
	zeroPtr(&i)
	fmt.Println(i)

	fmt.Println(&i)

}
