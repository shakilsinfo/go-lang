package main

import "fmt"

var balance float64 = 1000

func main() {
	for {
		var choice int
		fmt.Println("\n----- Go Bank -----")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			checkBalance()
		case 2:
			deposit()
		case 3:
			withdraw()
		case 4:
			fmt.Println("Thank you for using Go Bank. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func checkBalance() {
	fmt.Println("Your balance is:", balance)
}

func deposit() {
	var amount float64
	fmt.Print("Enter deposit amount: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Invalid amount!")
		return
	}

	balance += amount
	fmt.Println("Deposit successful! New balance:", balance)
}

func withdraw() {
	var amount float64
	fmt.Print("Enter withdrawal amount: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("Invalid amount!")
		return
	}

	if amount > balance {
		fmt.Println("Insufficient balance!")
		return
	}

	balance -= amount
	fmt.Println("Withdrawal successful! New balance:", balance)
}
