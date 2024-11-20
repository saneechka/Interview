package main

import "fmt"

func main() {
	// Creating a map using make
	ages := make(map[string]int)
	ages["Alice"] = 31
	ages["Bob"] = 42

	// Creating a map using a map literal
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println("Map of colors:", colors)

	// Adding key-value pairs
	ages["Charlie"] = 23

	// Deleting a key-value pair
	delete(ages, "Alice")

	// Checking if a key exists
	age, ok := ages["Bob"]
	if ok {
		fmt.Println("Bob's age:", age)
	}

	// Iterating over a map
	for name, age := range ages {
		fmt.Printf("%s is %d years old.\n", name, age)
	}

	// Map are reference types
	updateMap(ages)
	fmt.Println("Ages after update:", ages)
}

// Any changes to the map inside this function
// are visible to the caller
func updateMap(m map[string]int) {
	m["Bob"] = 50
}

