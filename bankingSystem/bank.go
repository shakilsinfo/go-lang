package main

import "fmt"

var accountBalance float64 = 1000.0 // Global account balance

func main() {
	var choice int
	fmt.Println("Welcome to the Banking System")
	fmt.Println("1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. Check Balance")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		checkBalance()
		return
	} else if choice == 2 {
		makeDeposit()
		return
	} else if choice == 3 {
		makeWithdrawal()
		return
	} else {
		fmt.Println("Invalid choice")
	}

}

func checkBalance() {
	fmt.Printf("Your account balance is: $%.2f\n", accountBalance)
}

func makeDeposit() {
	var amount float64
	fmt.Print("Enter amount to deposit: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Deposit amount must be positive")
		return
	}
	accountBalance += amount // accountBalance = accountBalance + amount
	fmt.Printf("Successfully deposited $%.2f. New balance is $%.2f\n", amount, accountBalance)
}

func makeWithdrawal() {
	var amount float64
	fmt.Print("Enter amount to withdraw: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be positive")
		return
	}
	if amount > accountBalance {
		fmt.Println("Insufficient funds for this withdrawal")
		return
	}
	accountBalance -= amount
	fmt.Printf("Successfully withdrew $%.2f. New balance is $%.2f\n", amount, accountBalance)
}
