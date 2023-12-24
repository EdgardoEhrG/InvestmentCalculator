package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate float64 = 2.5

	var (
		investmentAmount int = 1000
		expectedReturnRate float64 = 5.5
		years int = 10
	)

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := float64(investmentAmount) * math.Pow((1 + expectedReturnRate / 100), float64(years))
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, float64(years))

	fmt.Println("Future Value: ", futureValue)
	fmt.Println("Future Real Value: ", futureRealValue)
}