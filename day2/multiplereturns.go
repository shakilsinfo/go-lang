package main

import "fmt"

func divide(a int, b int) (int, string) {
	if b == 0 {
		return 0, "Division by zero error"
	}
	return a / b, "Success"
}
func main() {
	result, msg := divide(10, 2)
	fmt.Println("Result:", result, "Message:", msg)
	result, msg = divide(10, 15)
	fmt.Println("Result:", result, "Message:", msg)
}
