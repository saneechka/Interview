package main

import (
    "fmt"
    "log"
    "runtime/debug"
)

func main() {
    simplePanic()
    recoverFromPanic()
    panicInGoroutine()
    panicWithExtraInfo()
}

// simplePanic demonstrates a basic panic scenario.
func simplePanic() {
    fmt.Println("\n--- simplePanic ---")
    defer fmt.Println("Defer in simplePanic")
    fmt.Println("About to panic")
    panic("A simple panic happened!")
    // This line will not be executed.
    fmt.Println("After panic in simplePanic")
}

// recoverFromPanic demonstrates recovering from a panic.
func recoverFromPanic() {
    fmt.Println("\n--- recoverFromPanic ---")
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in recoverFromPanic:", r)
        }
    }()
    fmt.Println("About to panic in recoverFromPanic")
    panic("A panic from which we recover!")
    // This line will not be executed.
    fmt.Println("After panic in recoverFromPanic")
}

// panicInGoroutine demonstrates a panic inside a goroutine.
func panicInGoroutine() {
    fmt.Println("\n--- panicInGoroutine ---")
    go func() {
        defer fmt.Println("Defer in goroutine")
        fmt.Println("About to panic in a goroutine")
        panic("A panic inside a goroutine!")
    }()
    // Sleep is used here for demonstration purposes only.
    // In real-world applications, synchronization techniques like WaitGroup should be used.
    fmt.Println("A goroutine has panicked. Main continues.")
}

// panicWithExtraInfo demonstrates adding extra information to a panic.
func panicWithExtraInfo() {
    fmt.Println("\n--- panicWithExtraInfo ---")
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Recovered in panicWithExtraInfo: %v\n", r)
            log.Println("Stacktrace from panic:\n", string(debug.Stack()))
        }
    }()
    fmt.Println("About to panic with extra info")
    panic(fmt.Sprintf("A panic with additional info: %s", "important details"))
}


