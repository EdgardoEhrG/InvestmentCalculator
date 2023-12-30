package main

import (
	"fmt"
)

const inflationRate float64 = 2.5

func main() {

	var (
		investmentAmount,
		expectedReturnRate,
		years float64
	)

	// ----------------------- Input parameters

	printText("Investment Amount: $ ")
	fmt.Scan(&investmentAmount)

	printText("Expected Return Rate: % ")
	fmt.Scan(&expectedReturnRate)

	printText("Years: ")
	fmt.Scan(&years)

	// ----------------------- Calculation

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// ----------------------- Formatting

	formattedFV := fmt.Sprintf("Future Value: $ %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (AFI): $ %.1f\n", futureRealValue)

	// ----------------------- Outputs

	printOutputText(formattedFV, formattedRFV)
	printAditionalInfo()
}





