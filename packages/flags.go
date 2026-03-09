package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define command-line flags
	name := flag.String("name", "World", "a name to say hello to")
	age := flag.Int("age", 30, "your age")

	// Parse the flags
	flag.Parse()

	// Use the flag values
	fmt.Printf("Hello, %s! You are %d years old.\n", *name, *age)
}
