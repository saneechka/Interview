
package main

import (
    "fmt"
    "time"
)

func main() {
    // Example of an unbuffered channel
    unbufferedChan := make(chan int)

    go func() {
        fmt.Println("Sending value 1 to unbuffered channel")
        unbufferedChan <- 1 // Sending a value to the channel
        fmt.Println("Value 1 sent to unbuffered channel")
    }()

    // Receiving a value from the unbuffered channel
    value := <-unbufferedChan
    fmt.Println("Received from unbuffered channel:", value)

    // Example of a buffered channel
    bufferedChan := make(chan int, 2)

    // Sending values to a buffered channel
    bufferedChan <- 2
    bufferedChan <- 3

    // Receiving values from a buffered channel
    fmt.Println("Received from buffered channel:", <-bufferedChan)
    fmt.Println("Received from buffered channel:", <-bufferedChan)

    // Example of using select
    go func() {
        for i := 0; i < 2; i++ {
            time.Sleep(1 * time.Second)
            fmt.Println("Sending to channel:", i)
            bufferedChan <- i
        }
        close(bufferedChan) // Closing the channel after sending data
    }()

    // Receiving from the channel until it's closed
    for v := range bufferedChan {
        fmt.Println("Received from channel:", v)
    }

    fmt.Println("Channel closed, exiting main")
}

