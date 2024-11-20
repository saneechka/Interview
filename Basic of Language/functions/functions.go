package main

import "fmt"

// Simple function with no parameters and no return value
func sayHello() {
	fmt.Println("Hello, World!")
}

// Function with parameters
func greet(name string) {
	fmt.Println("Hello,", name)
}

// Function with return value
func add(x int, y int) int {
	return x + y
}

// Variadic function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Function as a parameter
func compute(fn func(int, int) int, a int, b int) int {
	return fn(a, b)
}

// Anonymous function and closure
func sequenceGenerator() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	sayHello()
	greet("Alice")

	result := add(3, 4)
	fmt.Println("3 + 4 =", result)

	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of 1,2,3,4,5 is", total)

	multiply := func(x, y int) int { return x * y }
	fmt.Println("3 * 4 =", compute(multiply, 3, 4))

	next := sequenceGenerator()
	fmt.Println(next()) // 1
	fmt.Println(next()) // 2
	fmt.Println(next()) // 3

	// Deferred function call
	defer fmt.Println("This will be printed last.")
	fmt.Println("This will be printed first.")
}

