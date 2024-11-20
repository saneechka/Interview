package main

import "fmt"

// Global variable
var globalVar = "I'm global"

// Exported struct with both exported and unexported fields
type ExportedStruct struct {
	ExportedField   string
	unexportedField string
}

// Unexported struct, only accessible within the same package
type unexportedStruct struct {
	Field string
}

func main() {
	// Local variable to main function
	localVar := "I'm local to main"

	{
		// Block scope variable
		blockVar := "I'm local to this block"
		fmt.Println(blockVar)
		fmt.Println(localVar)
	}

	// blockVar is not accessible here, uncommenting the next line would raise an error
	// fmt.println(blockVar)

	fmt.Println(localVar)
	accessGlobal()

	// Accessing the exported struct and its fields
	expStruct := ExportedStruct{"Visible to everyone", "Hidden from other packages"}
	fmt.Println(expStruct.ExportedField)
	// fmt.println(expStruct.unexportedField) // This line would raise an error if uncommented

	// Accessing the unexported struct within the same package
	unexpStruct := unexportedStruct{"I am only visible within the package"}
	fmt.Println(unexpStruct.Field)
}

func accessGlobal() {
	// access the global variable
	fmt.Println(globalVar)
}
