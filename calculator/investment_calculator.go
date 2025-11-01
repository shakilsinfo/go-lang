package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 3.5
	var investment float64 // Variable to hold investment amount as input
	expectedReturn := 5.5
	years := 10
	fmt.Scan(&investment) // Taking investment amount as input from user
	futureValue := float64(investment) * math.Pow((1+expectedReturn/100), float64(years))
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), float64(years))

	fmt.Printf("Future Value of the Investment: %.2f\n", futureValue)
	fmt.Println("Future Real Value of the Investment:", futureRealValue)
}
// Here in this example we showing variables, types and operators in Go program
// We are calculating future value of an investment using compound interest formula
