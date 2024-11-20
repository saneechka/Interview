package main

import "fmt"

func main() {
	// Creating a slice with make
	s1 := make([]int, 5) // slice of 5 integers
	fmt.Println("Slice with make:", s1)

	// Creating a slice from an array
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := array[1:5] // includes elements 1 through 4
	fmt.Println("Slice from array:", s2)

	// Creating a slice from another slice
	s3 := s2[:2] // includes elements 0 through 1 of s2
	fmt.Println("Slice from a slice:", s3)

	// Append to a slice
	s4 := append(s2, 10, 11, 12)
	fmt.Println("Slice after append:", s4)

	// Copy a slice
	s5 := make([]int, len(s4))
	copy(s5, s4)
	fmt.Println("Slice after copy:", s5)

	// Slice length and capacity
	fmt.Println("Length of s5:", len(s5))
	fmt.Println("Capacity of s5:", cap(s5))
}

