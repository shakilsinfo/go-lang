package main

import "fmt"

func add(a int, b int) int {
    sum := a + b
    return sum
}

func main() {
    result := add(5, 3)
    fmt.Println("Result:", result)
}