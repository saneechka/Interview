package main

import "fmt"

// Defining an interface
type Speaker interface {
  Speak() string
}

// Implementing the interface on a concrete type
type Dog struct {
  Name string
}

func (d Dog) Speak() string {
  return fmt.Sprintf("%s says: Woof!", d.Name)
}

// Implementing the interface on another concrete type
type Robot struct {
  Model string
}

func (r Robot) Speak() string {
  return fmt.Sprintf("Model %s says: Beep boop!", r.Model)
}

func performSpeak(s Speaker) {
  fmt.Println(s.Speak())
  
  // Type assertion
  if dog, ok := s.(Dog); ok {
    fmt.Println(dog.Name, "is a dog and can speak.")
  } else {
    fmt.Println("This Speaker is not a Dog.")
  }
}


// Defining simple interfaces
type Walker interface {
  Walk() string
}

type Talker interface {
  Talk() string
}

// Defining an embedded interface
type WalkerTalker interface {
  Walker
  Talker
}

// Implementing the simple interfaces on a concrete type
type Human struct {
  Name string
}

func (h Human) Walk() string {
  return fmt.Sprintf("%s is walking!", h.Name)
}

func (h Human) Talk() string {
  return fmt.Sprintf("%s says: Hello!", h.Name)
}

// Function that works with the embedded interface
func performAction(wt WalkerTalker) {
  fmt.Println(wt.Walk())
  fmt.Println(wt.Talk())
  
  // Type assertion
  if human, ok := wt.(Human); ok {
    fmt.Println(human.Name, "is a human and can walk and talk.")
  } else {
    fmt.Println("This WalkerTalker is not a Human.")
  }
}

func embeddedInterface() {
  john := Human{"John"}

  performAction(john)

  // Type assertions with the simple interfaces
  var wt WalkerTalker = john
  _, ok := wt.(Walker)
  fmt.Println("Is the WalkerTalker a Walker? ", ok)

  _, ok = wt.(Talker)
  fmt.Println("Is the WalkerTalker a Talker? ", ok)
}



// function that takes an empty interface as parameter
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func emptyInterface() {
	var any interface{}
	describe(any) // will print: (<nil>, <nil>)

	any = 42
	describe(any) // will print: (42, int)

	any = "hello"
	describe(any) // will print: (hello, string)

	any = struct{ name string }{"Alice"}
	describe(any) // will print: ({Alice}, struct { name string })
}


func main() {
  dog := Dog{"Buddy"}
  robot := Robot{"XJ-9"}

  performSpeak(dog)
  performSpeak(robot)

  // Let's try to assert to a different type
  var s Speaker = dog
  _, ok := s.(Dog)
  fmt.Println("Is the Speaker a Dog? ", ok)

  // This will be false
  _, ok = s.(Robot)
  fmt.Println("Is the Speaker a Robot? ", ok)

  embeddedInterface()
  emptyInterface()
}

