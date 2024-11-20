package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// Boolean type
	var isTrue bool = true
	fmt.Printf("Boolean: %t\n", isTrue)

	// String type
	var greeting string = "Hello, Go!"
	fmt.Printf("String: %s\n", greeting)

	// Integer types
	var age int = 30
	var temperature int64 = -25
	fmt.Printf("Integer: %d, Integer64: %d\n", age, temperature)

	// Unsigned integer types
	var distance uint = 4294967295
	fmt.Printf("Unsigned Integer: %d\n", distance)

	// Floating-point types
	var pi float64 = 3.14159
	var e float32 = 2.71828
	fmt.Printf("Float64: %f, Float32: %f\n", pi, e)

	// Complex number types
	var z complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Printf("Complex128: %v\n", z)

	// Array type
	var digits [5]int = [5]int{0, 1, 2, 3, 4}
	fmt.Printf("Array: %v\n", digits)

	// Slice type
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("Slice: %v\n", primes)

	// Map type
	var capitals map[string]string = map[string]string{"France": "Paris", "Italy": "Rome"}
	fmt.Printf("Map: %v\n", capitals)

	// Struct type
	type Point struct {
		X, Y int
	}
	var p Point = Point{1, 2}
	fmt.Printf("Struct: %+v\n", p)

	// Channel type
	ch := make(chan int, 1)
	ch <- 1
	fmt.Printf("Channel: %v\n", <-ch)
}

