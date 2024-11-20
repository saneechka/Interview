package main

import (
    "fmt"
    "strings"
)

// Map applies the function f to each item of the slice s and returns a new slice of results.
func Map[T any, U any](s []T, f func(T) U) []U {
    r := make([]U, len(s)) // Create a slice of type U with the same length as s
    for i, v := range s {  // Iterate over slice s
        r[i] = f(v) // Apply the function f to s[i] and store the result in r[i]
    }
    return r // Return the new slice with results
}

func main() {
    // Example with int slice
    ints := []int{1, 2, 3, 4}
    incInts := Map(ints, func(i int) int { // Increment each integer by 1
        return i + 1
    })
    fmt.Println(incInts) // Output: [2 3 4 5]

    // Example with string slice
    strs := []string{"a", "b", "c"}
    upperStrs := Map(strs, func(s string) string { // Convert each string to uppercase
        return strings.ToUpper(s)
    })
    fmt.Println(upperStrs) // Output: ["A", "B", "C"]
}

