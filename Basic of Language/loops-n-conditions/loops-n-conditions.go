package main

import "fmt"

func main() {
	// For loop as a "while" loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Traditional for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Infinite loop
	for {
		fmt.Println("loop")
		break
	}

	// If statement
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// Switch statement
	switch os := "linux"; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Nested loops
	for n := 0; n < 5; n++ {
		for m := 0; m <= n; m++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Using continue
	for x := 0; x < 10; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Println(x, "is odd")
	}

  // Range over a slice
  nums := []int{2, 3, 4}
  sum := 0
  for _, num := range nums {
      sum += num
  }
  fmt.Println("sum:", sum)

  // Range over a slice to get index and value
  for i, num := range nums {
      if num == 3 {
          fmt.Println("index:", i)
      }
  }

  // Range over a map
  kvs := map[string]string{"a": "apple", "b": "banana"}
  for k, v := range kvs {
      fmt.Printf("%s -> %s\n", k, v)
  }

  // Range over a string
  for i, c := range "go" {
      fmt.Println(i, c)
  }
}

