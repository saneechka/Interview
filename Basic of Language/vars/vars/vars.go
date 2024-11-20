package main

import "fmt"

func main() {
	// Declaring a variable with a specified type, which gets the zero value of that type
	var defaultInt int
	fmt.Println("Default value of int:", defaultInt)

	// Declaring a variable with an initial value
	var initializedInt int = 10
	fmt.Println("Initialized int:", initializedInt)

	// Short declaration of a variable, where the type is inferred from the value
	inferredInt := 20
	fmt.Println("Inferred int:", inferredInt)

	// Demonstrating the zero values for other types
	var defaultFloat float64
	var defaultString string
	var defaultBool bool
	fmt.Println("Default float:", defaultFloat)
	fmt.Println("Default string:", defaultString)
	fmt.Println("Default bool:", defaultBool)

	// Using the blank identifier to ignore an unwanted return value from a function
	_, err := someFunction()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Demonstrating the scope of a variable
	if true {
		var scopedVar = "I am visible only within this if statement"
		fmt.Println(scopedVar)
	}
	// fmt.println(scopedVar) // This would cause an error because scopedVar is not visible outside the if statement

	// Declaring multiple variables
	var (
		name string = "Go Programmer"
		age  int    = 25
		city string = "San Francisco"
	)
	fmt.Println("Name:", name, "Age:", age, "City:", city)
}

func someFunction() (int, error) {
	// Let's assume that the function does something and returns a result and an error
	// In this example, we're ignoring the result and just returning the zero value for the error
	return 0, nil
}
