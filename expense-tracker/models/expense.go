package models

// UserProfile stores user information
type UserProfile struct {
	Name   string
	Salary float64
}

// Expense represents a single expense entry
type Expense struct {
	Date        string
	Description string
	Amount      float64
}

// Custom error types for better error handling
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

// ValidateExpense checks if expense data is valid
func ValidateExpense(date, description string, amount float64) error {
	if date == "" {
		return ValidationError{Field: "Date", Message: "cannot be empty"}
	}
	if description == "" {
		return ValidationError{Field: "Description", Message: "cannot be empty"}
	}
	if amount <= 0 {
		return ValidationError{Field: "Amount", Message: "must be greater than 0"}
	}
	return nil
}

// ValidateProfile checks if profile data is valid
func ValidateProfile(name string, salary float64) error {
	if name == "" {
		return ValidationError{Field: "Name", Message: "cannot be empty"}
	}
	if salary <= 0 {
		return ValidationError{Field: "Salary", Message: "must be greater than 0"}
	}
	return nil
}

// FormatExpense converts expense to string format for file storage
func FormatExpense(expense Expense) string {
	return expense.Date + "," + expense.Description + "," + formatFloat(expense.Amount)
}

// FormatProfile converts profile to string format for file storage
func FormatProfile(profile UserProfile) string {
	return profile.Name + "\n" + formatFloat(profile.Salary)
}

// formatFloat formats a float64 to 2 decimal places
func formatFloat(value float64) string {
	// Simple formatting to 2 decimal places
	return formatFloatWithPrecision(value, 2)
}

// formatFloatWithPrecision formats float to specific decimal places
func formatFloatWithPrecision(value float64, precision int) string {
	// Multiply by 100, round, then divide by 100 for 2 decimal places
	multiplier := 100.0
	rounded := roundToDecimal(value, multiplier)

	// Convert to string with proper formatting
	// Using simple string building
	intPart := int(rounded)
	decPart := int((rounded - float64(intPart)) * multiplier + 0.5)

	// Handle decimal part
	if decPart >= 100 {
		decPart = 99
	}
	if decPart < 10 {
		return floatToString(intPart) + ".0" + intToString(decPart)
	}
	return floatToString(intPart) + "." + intToString(decPart)
}

// roundToDecimal rounds a float to specific decimal places
func roundToDecimal(value float64, multiplier float64) float64 {
	return float64(int(value*multiplier+0.5)) / multiplier
}

// Helper functions for number to string conversion
func floatToString(value int) string {
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

func intToString(value int) string {
	if value == 0 {
		return "0"
	}

	var result string
	for value > 0 {
		digit := value % 10
		result = string(rune('0'+digit)) + result
		value = value / 10
	}
	return result
}
