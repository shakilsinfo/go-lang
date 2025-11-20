package main

import (
	"fmt"
	"os" // Importing the os package to handle file operations
	"strconv" // Importing strconv to convert string to float64
)
const accountFile = "account.txt"
func main() {
	file, err := os.Create(accountFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	var balance float64
	fmt.Print("Enter balance: ")
	fmt.Scan(&balance)
	writeToFile(balance)

	fmt.Println("File written successfully.")
	balanceFromFile := readFromFile()
	fmt.Print(balanceFromFile)
}
func writeToFile(balance float64) {
	balanceStr := fmt.Sprintf("Current Balance: %.2f\n", balance)
	os.WriteFile(accountFile, []byte(balanceStr), 0644)

}

func readFromFile() float64 {
	data, _ := os.ReadFile(accountFile)
	balanceStr := string(data)
	balance, _ := strconv.ParseFloat(balanceStr, 64) // Extracting the balance value from the string
	return balance
}