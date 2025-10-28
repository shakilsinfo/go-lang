package main

import (
	"fmt"
	"math"
)

func main() {
	var investment = 1000
	var expectedReturn = 5.5
	var years = 10
	var futureValue = float64(investment) * math.Pow((1+expectedReturn/100), float64(years))
	fmt.Printf("Future Value of the Investment: %.2f\n", futureValue)
}
// Here in this example we showing variables, types and operators in Go program
// We are calculating future value of an investment using compound interest formula
