package main

import "fmt"

// Enumerated constants using iota
const (
    First = iota // 0
    Second       // 1
    Third        // 2
)

// Typed constant declaration
const Pi float64 = 3.14159

// Untyped constant declaration
const Truth = true

func main() {
    // Using enumerated constants
    fmt.Println("First constant:", First)
    fmt.Println("Second constant:", Second)
    fmt.Println("Third constant:", Third)

    // Using typed constant
    fmt.Println("Value of Pi:", Pi)

    // Using untyped constant
    fmt.Println("Value of Truth:", Truth)

    // Constants can be used in expressions
    const Degree = 180 / Pi
    fmt.Println("Degree constant:", Degree)

    // Untyped constants can be used without type conversion
    var radius float64 = 5
    var circumference = 2 * Pi * radius
    fmt.Println("Circumference:", circumference)
}

// Constants can also be declared in a block
const (
    Monday    = iota + 1 // 1
    Tuesday              // 2
    Wednesday            // 3
    Thursday             // 4
    Friday               // 5
    Saturday             // 6
    Sunday               // 7
)

func printDay(day int) {
    fmt.Println("Day of the week:", day)
}

