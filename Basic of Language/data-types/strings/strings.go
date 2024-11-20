package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// String initialization
	hello := "Hello"
	world := "World"

	// Concatenation
	greeting := hello + ", " + world + "!"
	fmt.Println(greeting)

	// Multiline string
	multiline := `This is a
  multiline string.
  Special characters like \n do not need to be escaped.`
	fmt.Println(multiline)

	// Length of a string
	length := len(greeting)
	fmt.Println("Length:", length)

	// Accessing individual runes (characters) using range
	for index, runeValue := range greeting {
		fmt.Printf("%d -> %c\n", index, runeValue)
	}

	// Searching for a substring
	fmt.Println(strings.Contains(greeting, "Hello")) // true
	fmt.Println(strings.Index(greeting, "lo"))       // 3

	// Replacing substrings
	newGreeting := strings.Replace(greeting, "Hello", "Hi", -1)
	fmt.Println(newGreeting)

	// Converting integer to string
	num := 42
	strNum := strconv.Itoa(num)
	fmt.Println("String from integer:", strNum)

	// Converting string to integer
	if num, err := strconv.Atoi(strNum); err == nil {
		fmt.Println("Integer from string:", num)
	}
}

