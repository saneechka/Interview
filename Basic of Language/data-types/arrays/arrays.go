package main

import "fmt"

func main() {
	// Declaring an array with a specified size
	var a [5]int
	fmt.Println("Empty array:", a)

	// Initializing an array with elements
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Initialized array:", b)

	// Letting Go infer the array size
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array with inferred size:", c)

	// Accessing elements
	fmt.Println("First element of b:", b[0])
	fmt.Println("Second element of c:", c[1])

	// Iterating over an array with a for loop
	for i := 0; i < len(b); i++ {
		fmt.Printf("b[%d]: %d\n", i, b[i])
	}

	// Iterating over an array with a for range loop
	for index, value := range c {
		fmt.Printf("c[%d]: %d\n", index, value)
	}

	// Changing an element
	b[2] = 100
	fmt.Println("Modified array b:", b)
}


