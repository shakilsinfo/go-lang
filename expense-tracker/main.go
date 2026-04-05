package main

import (
	"expense-tracker/services"
	"fmt"
)

func main() {
	fmt.Println("=================================")
	fmt.Println("    EXPENSE TRACKER v1.0")
	fmt.Println("=================================")

	// Check if user has a profile
	hasProfile := services.HasProfile()
	if !hasProfile {
		fmt.Println("\nWelcome! It looks like you're new here.")
		fmt.Println("Let's set up your profile first.")

		err := services.SetupProfile()
		if err != nil {
			fmt.Printf("Error setting up profile: %v\n", err)
			return
		}
	} else {
		// Load and show profile info
		name, salary, totalExpense, currentCash, err := services.GetCurrentCash()
		if err != nil {
			fmt.Printf("Error loading profile: %v\n", err)
			return
		}
		fmt.Printf("\nWelcome back, %s!\n", name)
		fmt.Printf("Salary: %.2f | Expenses: %.2f | Cash in Hand: %.2f\n", salary, totalExpense, currentCash)
	}

	// Main menu loop
	for {
		fmt.Println("\n=================================")
		fmt.Println("           MAIN MENU")
		fmt.Println("=================================")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Current Cash")
		fmt.Println("3. View All Expenses")
		fmt.Println("4. View Expense Summary")
		fmt.Println("5. Exit")
		fmt.Println("=================================")
		fmt.Print("Enter your choice (1-5): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			err := services.AddExpense()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case 2:
			err := services.ShowCurrentCash()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case 3:
			err := services.ShowAllExpenses()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case 4:
			err := services.ShowExpenseSummary()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case 5:
			fmt.Println("\nThank you for using Expense Tracker!")
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("\nInvalid choice! Please enter a number between 1 and 5.")
		}
	}
}
