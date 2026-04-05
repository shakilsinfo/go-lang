package main

import (
	"fmt"
	"math"
)

// ==================== CUSTOM ERROR TYPES ====================

// DivisionError is a custom error type for division-related errors
type DivisionError struct {
	Dividend float64
	Divisor  float64
	Message  string
}

// Error implements the error interface
func (e *DivisionError) Error() string {
	return fmt.Sprintf("%s: cannot divide %.2f by %.2f", e.Message, e.Dividend, e.Divisor)
}

// InsufficientFundsError is another custom error type
type InsufficientFundsError struct {
	Required float64
	Available float64
	Account string
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("account %s: insufficient funds (required: $%.2f, available: $%.2f)",
		e.Account, e.Required, e.Available)
}

// ==================== FUNCTIONS THAT RETURN CUSTOM ERRORS ====================

// divideWithCustomError returns a custom DivisionError for division by zero
func divideWithCustomError(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, &DivisionError{
			Dividend: dividend,
			Divisor:  divisor,
			Message:  "division by zero",
		}
	}
	return dividend / divisor, nil
}

// calculateSqrt returns an error for negative numbers
func calculateSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &DivisionError{
			Dividend: x,
			Divisor:  0,
			Message:  "cannot calculate square root of negative number",
		}
	}
	return math.Sqrt(x), nil
}

// withdraw returns InsufficientFundsError when balance is too low
func withdraw(balance, amount float64, account string) (float64, error) {
	if amount > balance {
		return 0, &InsufficientFundsError{
			Required:  amount,
			Available: balance,
			Account:   account,
		}
	}
	return balance - amount, nil
}

// ==================== PANIC AND RECOVER EXAMPLES ====================

// criticalOperation demonstrates panic usage
func criticalOperation(value int) {
	if value < 0 {
		panic(fmt.Sprintf("critical failure: negative value %d not allowed", value))
	}
	fmt.Printf("Operation completed successfully with value: %d\n", value)
}

// safeOperation demonstrates panic recovery with defer
func safeOperation(value int) (recovered bool) {
	// defer must be called before the panic can happen
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
			recovered = true
		}
	}()

	if value < 0 {
		panic(fmt.Sprintf("panic triggered: value %d is invalid", value))
	}
	fmt.Printf("Operation completed with value: %d\n", value)
	return recovered
}

// processItemWithPanic demonstrates panic in a slice processing context
func processItemWithPanic(items []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic recovered while processing item at index %d: %v\n", index, r)
		}
	}()

	if index >= len(items) {
		panic(fmt.Sprintf("index out of bounds: %d (len=%d)", index, len(items)))
	}

	if items[index] < 0 {
		panic(fmt.Sprintf("invalid item value: %d at index %d", items[index], index))
	}

	fmt.Printf("Processed item: %d\n", items[index])
}

// ==================== MAIN FUNCTION ====================

func main() {
	fmt.Println("========== CUSTOM ERROR HANDLING EXAMPLES ==========\n")

	// Example 1: Division by zero with custom error
	fmt.Println("1. Division Error:")
	result, err := divideWithCustomError(10, 0)
	if err != nil {
		fmt.Println("  Error:", err)
		// Type assertion to access custom error fields
		if divErr, ok := err.(*DivisionError); ok {
			fmt.Printf("  Details - Dividend: %.2f, Divisor: %.2f\n", divErr.Dividend, divErr.Divisor)
		}
	} else {
		fmt.Println("  Result:", result)
	}

	// Example 2: Square root of negative number
	fmt.Println("\n2. Square Root Error:")
	sqrtResult, err := calculateSqrt(-4)
	if err != nil {
		fmt.Println("  Error:", err)
	} else {
		fmt.Println("  Result:", sqrtResult)
	}

	// Example 3: Insufficient funds error
	fmt.Println("\n3. Insufficient Funds Error:")
	newBalance, err := withdraw(100.00, 150.00, "ACC-1234")
	if err != nil {
		fmt.Println("  Error:", err)
	} else {
		fmt.Printf("  New balance: $%.2f\n", newBalance)
	}

	fmt.Println("\n========== PANIC AND RECOVER EXAMPLES ==========\n")

	// Example 4: Unhandled panic (will crash if not recovered)
	fmt.Println("4. Safe Operation with Panic Recovery:")
	recovered := safeOperation(-5)
	fmt.Printf("  Was panic recovered? %v\n", recovered)

	fmt.Println("\n5. Safe Operation with Valid Value:")
	recovered = safeOperation(10)
	fmt.Printf("  Was panic recovered? %v\n", recovered)

	// Example 6: Processing items with potential panic
	fmt.Println("\n6. Processing Items with Panic Recovery:")
	items := []int{10, 20, -5, 40, 50}
	for i := 0; i <= len(items); i++ {
		fmt.Printf("  Processing index %d: ", i)
		processItemWithPanic(items, i)
	}

	fmt.Println("\n========== COMPARISON: ERROR VS PANIC ==========")
	fmt.Println("\nWhen to use errors:")
	fmt.Println("  - Predictable problems (file not found, invalid input)")
	fmt.Println("  - Situations the caller can handle")
	fmt.Println("  - External factors (network issues, user input)")

	fmt.Println("\nWhen to use panic:")
	fmt.Println("  - Truly unrecoverable conditions")
	fmt.Println("  - Programming errors (should never happen in correct code)")
	fmt.Println("  - Initialization failures that prevent the program from running")

	fmt.Println("\nAlways use defer + recover in production to handle panics gracefully!")
}