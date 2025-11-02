package main
import "fmt"
func main() {
	// We are trying to build a simple profit calculator
	var revenue, expenses, taxRate float64
	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter Tax Rate (in %): ")
	fmt.Scan(&taxRate)

	profitBeforeTax := revenue - expenses
	profit := profitBeforeTax * (1- taxRate / 100)
	profitMargin := profitBeforeTax / profit

	if(revenue < expenses){
		fmt.Println("Uhhppsss!, the business is running at a loss. No worries. Best time to rethink strategy and improve!")
		fmt.Printf("Total loss amount including tax: %.2f\n", profit)
	}else{
		fmt.Printf("Profit exclude Tax: %.2f\n", profitBeforeTax)
		fmt.Printf("Profit After Tax: %.2f\n", profit)
		fmt.Printf("Profit Margin: %.2f\n", profitMargin)
	}

}	