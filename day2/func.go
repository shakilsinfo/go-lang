package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	greet("Shakel")
	greet("Ahamed")
}
