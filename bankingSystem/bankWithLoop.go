package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration %d: Welcome to the Banking System\n", i)
		var choice int
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			checkBalance1()
			continue // Skip to the next iteration
		} else if choice == 2 {
			makeDeposit1()
		} else if choice == 3 {
			makeWithdrawal1()
		} else {
			fmt.Println("Invalid choice")
			break // Exit the loop on invalid choice
		}
		fmt.Println() // Print a newline for better readability between iterations
	}
	fmt.Println("Thanks for choosing our banking system! Goodbye!")
}

func checkBalance1() {
	fmt.Println("Checking balance...")
}

func makeDeposit1() {
	fmt.Println("Making deposit...")
}

func makeWithdrawal1() {
	fmt.Println("Making withdrawal...")
}
