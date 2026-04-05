package storage

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// File constants
const (
	ProfileFileName  = "profile.txt"
	ExpensesFileName = "expenses.txt"
)

// ProfileExists checks if profile file exists
func ProfileExists() bool {
	_, err := os.Stat(ProfileFileName)
	return err == nil
}

// SaveProfile saves user profile to file
func SaveProfile(name string, salary float64) error {
	content := name + "\n" + formatFloat(salary) + "\n"
	err := os.WriteFile(ProfileFileName, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to save profile: %w", err)
	}
	return nil
}

// LoadProfile reads user profile from file
func LoadProfile() (string, float64, error) {
	content, err := os.ReadFile(ProfileFileName)
	if err != nil {
		return "", 0, fmt.Errorf("failed to read profile: %w", err)
	}

	lines := strings.Split(string(content), "\n")
	if len(lines) < 2 {
		return "", 0, errors.New("invalid profile file format")
	}

	name := strings.TrimSpace(lines[0])
	salaryStr := strings.TrimSpace(lines[1])
	salary := parseFloat(salaryStr)

	if name == "" {
		return "", 0, errors.New("profile name is empty")
	}

	return name, salary, nil
}

// SaveExpense appends a single expense to the expenses file
func SaveExpense(date, description string, amount float64) error {
	file, err := os.OpenFile(ExpensesFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open expenses file: %w", err)
	}
	defer file.Close()

	line := date + "," + description + "," + formatFloat(amount) + "\n"
	_, err = file.WriteString(line)
	if err != nil {
		return fmt.Errorf("failed to write expense: %w", err)
	}

	return nil
}

// LoadExpenses reads all expenses from file
func LoadExpenses() ([]map[string]string, error) {
	file, err := os.Open(ExpensesFileName)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist yet, return empty slice
			return []map[string]string{}, nil
		}
		return nil, fmt.Errorf("failed to open expenses file: %w", err)
	}
	defer file.Close()

	var expenses []map[string]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) >= 3 {
			expense := map[string]string{
				"date":        parts[0],
				"description": parts[1],
				"amount":      parts[2],
			}
			expenses = append(expenses, expense)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading expenses: %w", err)
	}

	return expenses, nil
}

// Helper function to format float to string with 2 decimal places
func formatFloat(value float64) string {
	// Round to 2 decimal places
	multiplier := 100.0
	rounded := float64(int(value*multiplier+0.5)) / multiplier

	// Convert to string
	intPart := int(rounded)
	decPart := int((rounded - float64(intPart)) * multiplier + 0.5)

	// Handle negative numbers
	isNegative := rounded < 0
	if isNegative {
		intPart = -intPart
	}

	// Build string
	result := intToString(intPart)
	result = result + "."

	if decPart < 10 {
		result = result + "0" + intToString(decPart)
	} else {
		result = result + intToString(decPart)
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

// Helper function to parse float from string
func parseFloat(s string) float64 {
	// Simple implementation - handles positive decimals
	result := 0.0
	decimalPlaces := 0.0
	afterDecimal := false
	isNegative := false

	for i, char := range s {
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

// Helper function to convert int to string
func intToString(value int) string {
	if value == 0 {
		return "0"
	}

	isNegative := value < 0
	if isNegative {
		value = -value
	}

	var result string
	for value > 0 {
		digit := value % 10
		result = string(rune('0'+digit)) + result
		value = value / 10
	}

	if isNegative {
		result = "-" + result
	}
	return result
}
