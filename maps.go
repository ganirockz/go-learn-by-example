package main

import "fmt"

func main() {

	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	m["k1"] = 7

	m["k2"] = 13

	fmt.Println(m)

	v1 := m["k1"]

	fmt.Println("v1:", v1)

	// If the key doesn’t exist, the zero value of the value type is returned.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map
	delete(m, "k2")

	fmt.Println(m)

	// we will get two values for a key. first is value and second is a boolean whether value is available or not. Since we don't value we can '_' to ignore the value.
	_, prs := m["k2"]
	fmt.Println("Prs:", prs)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
