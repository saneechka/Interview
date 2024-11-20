package main

import "fmt"

type Counter struct {
    Value int
}

func IncrementPointer(c *Counter) {
    c.Value++
}

func IncrementValue(c Counter) {
    c.Value++
}

func main() {

    counterPointer := &Counter{Value: 0}
    fmt.Println("Pointer before:", counterPointer.Value)
    IncrementPointer(counterPointer)
    fmt.Println("Pointer after:", counterPointer.Value) 

    counterValue := Counter{Value: 0}
    fmt.Println("Value before:", counterValue.Value)
    IncrementValue(counterValue)
    fmt.Println("Value after:", counterValue.Value) }

