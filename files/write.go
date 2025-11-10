package main

import (
	"fmt"
	"os" // Importing the os package to handle file operations
)

func main() {
	file, err := os.Create("output.txt")
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
}
func writeToFile(balance float64) {
	balanceStr := fmt.Sprintf("Current Balance: %.2f\n", balance)
	os.WriteFile("balance.txt", []byte(balanceStr), 0644)

}
