package main

import "fmt"

// Define a new struct type
type Person struct {
	Name string
	Age  int
}

// Function to create a new person
func NewPerson(name string, age int) *Person {
	p := Person{Name: name, Age: age}
	return &p
}

// Using new() to create a pointer to a new struct instance
func NewPersonWithNew(name string, age int) *Person {
	p := new(Person)
	p.Name = name
	p.Age = age
	return p
}

// Method for the Person struct
func (p *Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name, "and I am", p.Age, "years old.")
}

// Employee now embeds a pointer to Person
type Employee struct {
	*Person
	Company string
}

func main() {
	// Create a new Person instance
	john := NewPerson("John", 30)
	john.SayHello()

	// Accessing fields
	john.Age = 31 // John had a birthday!
	fmt.Println("John's new age is:", john.Age)

	// Anonymous struct
	tempPerson := struct {
		Name string
		Age  int
	}{
		Name: "Alice",
		Age:  25,
	}
	fmt.Println("Temporary person:", tempPerson)

	// Nested struct
	// type Employee struct {
	// 	Person  // Embedded struct
	// 	Company string
	// }
	// alice := Employee{
	// 	Person:  Person{Name: "Alice", Age: 28},
	// 	Company: "Tech Corp",
	// }

	// fmt.Printf("%s works at %s.\n", alice.Name, alice.Company)

  // Create a new Person instance with new()
	jane := NewPersonWithNew("Jane", 29)
	fmt.Println(jane)

	// Embedding a pointer to Person within Employee
	alice := &Employee{
		Person:  &Person{Name: "Alice", Age: 28},
		Company: "Tech Corp",
	}

	fmt.Printf("%s works at %s and is %d years old.\n", alice.Name, alice.Company, alice.Age)
}

