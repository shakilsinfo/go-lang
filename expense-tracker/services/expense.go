package services

import (
	"expense-tracker/models"
	"expense-tracker/storage"
	"fmt"
	"os"
)

// SetupProfile creates a new user profile
func SetupProfile() error {
	fmt.Println("\n--- Setup Your Profile ---")
	fmt.Print("Enter your name: ")

	var name string
	fmt.Scanln(&name)

	fmt.Print("Enter your monthly salary: ")
	var salary float64
	fmt.Scan(&salary)

	// Validate profile
	err := models.ValidateProfile(name, salary)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// Save profile
	err = storage.SaveProfile(name, salary)
	if err != nil {
		return fmt.Errorf("failed to save profile: %w", err)
	}

	fmt.Println("\n✓ Profile saved successfully!")
	return nil
}

// AddExpense adds a new expense entry
func AddExpense() error {
	fmt.Println("\n--- Add New Expense ---")

	fmt.Print("Enter date (DD-MM-YYYY): ")
	var date string
	fmt.Scanln(&date)

	fmt.Print("Enter description: ")
	var description string
	fmt.Scanln(&description)

	fmt.Print("Enter amount: ")
	var amount float64
	fmt.Scanln(&amount)

	// Validate expense
	err := models.ValidateExpense(date, description, amount)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// Save expense
	err = storage.SaveExpense(date, description, amount)
	if err != nil {
		return fmt.Errorf("failed to save expense: %w", err)
	}

	fmt.Printf("\n✓ Expense added: %s - %.2f\n", description, amount)
	return nil
}

// GetCurrentCash calculates and returns the current cash in hand
func GetCurrentCash() (string, float64, float64, float64, error) {
	// Load profile
	name, salary, err := storage.LoadProfile()
	if err != nil {
		return "", 0, 0, 0, fmt.Errorf("failed to load profile: %w", err)
	}

	// Load expenses
	expenses, err := storage.LoadExpenses()
	if err != nil {
		return "", 0, 0, 0, fmt.Errorf("failed to load expenses: %w", err)
	}

	// Calculate total expenses
	totalExpense := 0.0
	for i := 0; i < len(expenses); i++ {
		amountStr := expenses[i]["amount"]
		amount := parseAmount(amountStr)
		totalExpense = totalExpense + amount
	}

	// Calculate current cash
	currentCash := salary - totalExpense

	return name, salary, totalExpense, currentCash, nil
}

// GetAllExpenses returns all expenses
func GetAllExpenses() ([]map[string]string, error) {
	expenses, err := storage.LoadExpenses()
	if err != nil {
		return nil, fmt.Errorf("failed to load expenses: %w", err)
	}
	return expenses, nil
}

// parseAmount converts amount string to float64
func parseAmount(s string) float64 {
	// Simple implementation - handles positive decimals
	result := 0.0
	decimalPlaces := 0.0
	afterDecimal := false
	isNegative := false

	for i := 0; i < len(s); i++ {
		char := rune(s[i])

		if i == 0 && char == '-' {
			isNegative = true
			continue
		}

		if char == '.' {
			afterDecimal = true
			continue
		}

		digit := float64(char - '0')
		if afterDecimal {
			decimalPlaces = decimalPlaces*10 + digit
		} else {
			result = result*10 + digit
		}
	}

	// Add decimal part
	if decimalPlaces > 0 {
		// Count decimal digits
		temp := int(decimalPlaces)
		digits := 0
		for temp > 0 {
			temp = temp / 10
			digits++
		}

		divider := 1.0
		for i := 0; i < digits; i++ {
			divider = divider * 10
		}
		result = result + decimalPlaces/divider
	}

	if isNegative {
		result = -result
	}

	return result
}

// HasProfile checks if user has set up their profile
func HasProfile() bool {
	return storage.ProfileExists()
}

// DeleteProfile removes the profile file (for testing/reset)
func DeleteProfile() error {
	err := os.Remove(storage.ProfileFileName)
	if err != nil {
		return fmt.Errorf("failed to delete profile: %w", err)
	}
	return nil
}

// DeleteAllExpenses removes all expenses (for testing/reset)
func DeleteAllExpenses() error {
	err := os.Remove(storage.ExpensesFileName)
	if err != nil {
		return fmt.Errorf("failed to delete expenses: %w", err)
	}
	return nil
}
