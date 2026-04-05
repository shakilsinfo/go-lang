# Expense Tracker CLI Application

A simple command-line expense tracker built with Go, using only basic Go concepts (variables, functions, control flow, file I/O, error handling, structs, and slices).

## Project Structure

```
expense-tracker/
├── go.mod              # Go module definition
├── main.go             # Entry point with CLI menu
├── models/
│   └── expense.go      # Data models and validation
├── storage/
│   └── file.go         # File I/O operations (text format, NOT JSON)
├── services/
│   ├── expense.go      # Business logic for expenses
│   └── report.go       # Reporting functionality
├── profile.txt         # Created automatically - stores user profile
└── expenses.txt        # Created automatically - stores expense records
```

## How to Run

1. Navigate to the expense-tracker directory:
   ```bash
   cd /Users/shakelahamed/WorkStations/go-lang/expense-tracker
   ```

2. Build the application:
   ```bash
   go build
   ```

3. Run the application:
   ```bash
   ./expense-tracker
   ```

Or run directly:
```bash
go run main.go
```

## First Time Setup

When you run the application for the first time, it will prompt you to:
1. Enter your name
2. Enter your monthly salary

This information is saved in `profile.txt`

## Features

### 1. Add Expense
- Enter date in DD-MM-YYYY format
- Enter a description for the expense
- Enter the amount spent
- All expenses are saved in `expenses.txt`

### 2. View Current Cash
- Shows your current cash in hand (Salary - Total Expenses)
- Displays your salary and total expenses
- Warns if you have less than 20% of your salary remaining

### 3. View All Expenses
- Lists all recorded expenses with date, description, and amount
- Shows the total expenses at the bottom

### 4. View Expense Summary
- Comprehensive financial report
- Shows salary, total expenses, and remaining cash
- Displays expense ratio (percentage of salary spent)
- Shows spending status (Within budget, Moderate, High, or Over budget)
- Displays transaction count and average expense
- Highlights your largest expense

## File Format

### profile.txt
```
John Doe
50000.00
```

### expenses.txt
```
09-03-2026,Grocery Shopping,2500.50
10-03-2026,Electricity Bill,1500.00
11-03-2026,Transport,800.00
```

## Reset Data

To start fresh, delete the data files:
```bash
rm profile.txt expenses.txt
```

## Go Concepts Used

- **Variables**: var, :=, basic types (int, float64, string), constants
- **Functions**: parameters, multiple return values, named returns
- **Control Flow**: if/else, switch, for loops
- **File I/O**: os.Create, os.WriteFile, os.ReadFile, os.OpenFile, defer file.Close()
- **Error Handling**: if err != nil, errors.New(), custom error structs
- **User Input**: fmt.Scan, fmt.Scanln
- **Math**: basic operations (+, -, *, /), rounding
- **Slices**: basic append/len
- **Structs**: for data organization
- **fmt**: Printf, Sprintf for formatting

## Custom Error Types

The application uses custom validation errors:
```go
type ValidationError struct {
    Field   string
    Message string
}
```

This provides clear error messages when validating user input.
