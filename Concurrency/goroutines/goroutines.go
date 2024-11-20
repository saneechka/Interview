package main

import (
	"fmt"
	"sync"
	"time"
)

// A simple function that a goroutine will run
func printCounts(label string, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%s: %d\n", label, i)
	}
}

func main() {
	var wg sync.WaitGroup

	// Start a goroutine for printCounts
	wg.Add(1)
	go printCounts("goroutine", 5, &wg)

	// Run printCounts in the main goroutine as well
	printCounts("main", 3, &wg)

	// Wait for the goroutine to finish
	wg.Wait()
	fmt.Println("Finished all goroutines.")
}

