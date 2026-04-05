# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a learning-oriented Go repository where concepts are organized by "days" and topic areas. Each directory represents a separate Go module focusing on specific Go concepts (functions, packages, file I/O, error handling, financial calculations, etc.).

## Module Structure

Each top-level directory (day1/, day2/, day3/, calculator/, bankingSystem/, errors/, files/) is its own independent Go module with its own `go.mod` file. There is no monolithic structure.

**Important:** When working on any module, you must `cd` into that directory first before running Go commands.

## Building and Running

Since each directory is a separate module, commands must be run from within the specific module directory:

```bash
# Example: Running the banking system
cd bankingSystem
go run bankWithInfiniteLoop.go

# Example: Running the investment calculator
cd calculator
go run investment_calculator.go
```

For modules without a `go.mod` (e.g., simple single-file scripts in day1/, day2/):
```bash
go run day1/hello-world.go
```

## Common Commands

- **Run a Go file**: `go run <filename>.go`
- **Build a module**: `cd <module-dir> && go build`
- **Tidy dependencies**: `cd <module-dir> && go mod tidy`
- **Format code**: `go fmt ./...`

## Architecture Notes

- The codebase demonstrates progression from basic concepts (day1/day2) to more complex applications (bankingSystem, calculator)
- Financial calculations are present in calculator/ and profitCalculator/ - these use standard math operations
- Error handling varies across modules - from basic `if err != nil` checks to custom error types in errors/
- The bankingSystem/ module has multiple versions showing iterative development

## Testing

No test files currently exist in the codebase. When adding tests, place `*_test.go` files alongside the code they test in the appropriate module directory.
