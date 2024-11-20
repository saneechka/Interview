
package main

import (
    "errors"
    "fmt"
)

// Custom error type
type MyError struct {
    Msg  string
    Code int
}

// Implementing the error interface for MyError
func (e *MyError) Error() string {
    return fmt.Sprintf("code %d: %s", e.Code, e.Msg)
}

// Function that may return an error
func doSomething() (int, error) {
    // An error occurred
    return 0, &MyError{"Something went wrong", 123}
}

// Main function to demonstrate error handling
func main() {
    _, err := doSomething()
    if err != nil {
        var myErr *MyError
        if errors.As(err, &myErr) {
            fmt.Println("Custom error occurred:", myErr)
        } else {
            fmt.Println("An error occurred:", err)
        }
    }
}
