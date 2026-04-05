package services

import (
	"expense-tracker/storage"
	"fmt"
	"os"
)

// ShowCurrentCash displays the current cash in hand
func ShowCurrentCash() error {
	name, salary, totalExpense, currentCash, err := GetCurrentCash()
	if err != nil {
		return err
	}

	fmt.Println("\n=================================")
	fmt.Println("         CASH SUMMARY")
	fmt.Println("=================================")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Salary: %.2f\n", salary)
	fmt.Printf("Total Expenses: %.2f\n", totalExpense)
	fmt.Println("---------------------------------")
	fmt.Printf("Cash in Hand: %.2f\n", currentCash)
	fmt.Println("=================================")

	// Show warning if cash is low
	if currentCash < (salary * 0.2) {
		fmt.Println("\n⚠ Warning: You have less than 20% of your salary remaining!")
	}

	return nil
}

// ShowAllExpenses displays all expense entries
func ShowAllExpenses() error {
	expenses, err := GetAllExpenses()
	if err != nil {
		return err
	}

	fmt.Println("\n=================================")
	fmt.Println("         ALL EXPENSES")
	fmt.Println("=================================")

	if len(expenses) == 0 {
		fmt.Println("No expenses recorded yet.")
		fmt.Println("=================================")
		return nil
	}

	// Print header
	fmt.Printf("%-12s %-20s %10s\n", "Date", "Description", "Amount")
	fmt.Println("---------------------------------")

	// Print each expense
	for i := 0; i < len(expenses); i++ {
		expense := expenses[i]
		date := expense["date"]
		description := expense["description"]
		amountStr := expense["amount"]
		amount := parseAmount(amountStr)

		fmt.Printf("%-12s %-20s %10.2f\n", date, truncateString(description, 20), amount)
	}

	fmt.Println("=================================")

	// Show total
	totalExpense := 0.0
	for i := 0; i < len(expenses); i++ {
		amountStr := expenses[i]["amount"]
		amount := parseAmount(amountStr)
		totalExpense = totalExpense + amount
	}
	fmt.Printf("Total: %.2f\n", totalExpense)
	fmt.Println("=================================")

	return nil
}

// ShowExpenseSummary displays a summary of expenses
func ShowExpenseSummary() error {
	name, salary, totalExpense, currentCash, err := GetCurrentCash()
	if err != nil {
		return err
	}

	expenses, err := GetAllExpenses()
	if err != nil {
		return err
	}

	fmt.Println("\n=================================")
	fmt.Println("      EXPENSE SUMMARY REPORT")
	fmt.Println("=================================")
	fmt.Printf("User: %s\n\n", name)

	// Salary and expense info
	fmt.Printf("Monthly Salary: %.2f\n", salary)
	fmt.Printf("Total Expenses: %.2f\n", totalExpense)
	fmt.Printf("Remaining Cash: %.2f\n", currentCash)

	// Calculate percentage
	if salary > 0 {
		expensePercent := (totalExpense / salary) * 100
		fmt.Printf("Expense Ratio: %.1f%% of salary\n", expensePercent)

		// Show status
		fmt.Print("\nStatus: ")
		if expensePercent > 100 {
			fmt.Println("OVER BUDGET!")
		} else if expensePercent > 80 {
			fmt.Println("Warning - High expenses")
		} else if expensePercent > 50 {
			fmt.Println("Moderate spending")
		} else {
			fmt.Println("Good - Within budget")
		}
	}

	// Expense count
	fmt.Printf("\nTotal Transactions: %d\n", len(expenses))

	if len(expenses) > 0 {
		// Calculate average expense
		avgExpense := totalExpense / float64(len(expenses))
		fmt.Printf("Average Expense: %.2f\n", avgExpense)

		// Find largest expense
		maxAmount := 0.0
		maxIndex := 0
		for i := 0; i < len(expenses); i++ {
			amountStr := expenses[i]["amount"]
			amount := parseAmount(amountStr)
			if amount > maxAmount {
				maxAmount = amount
				maxIndex = i
			}
		}

		fmt.Printf("Largest Expense: %.2f (%s)\n", maxAmount, expenses[maxIndex]["description"])
	}

	fmt.Println("=================================")

	return nil
}

// truncateString truncates a string to specified length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[0:maxLen-3] + "..."
}

// ResetApplication deletes all data (use with caution)
func ResetApplication() error {
	fmt.Println("\n⚠ WARNING: This will delete ALL your data!")
	fmt.Print("Are you sure? (type 'yes' to confirm): ")

	var confirmation string
	fmt.Scanln(&confirmation)

	if confirmation != "yes" {
		fmt.Println("Operation cancelled.")
		return nil
	}

	// Delete expenses file
	err := deleteFileIfExists(storage.ExpensesFileName)
	if err != nil {
		return fmt.Errorf("failed to delete expenses: %w", err)
	}

	// Delete profile file
	err = deleteFileIfExists(storage.ProfileFileName)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}

	fmt.Println("✓ All data has been reset.")
	fmt.Println("Please restart the application to set up a new profile.")
	os.Exit(0)
	return nil
}

// deleteFileIfExists deletes a file if it exists
func deleteFileIfExists(filename string) error {
	_, err := os.Stat(filename)
	if err == nil {
		// File exists, delete it
		err := os.Remove(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
