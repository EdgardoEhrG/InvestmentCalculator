package main

import (
	"fmt"
	"math"
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

func printText (text string) {
	fmt.Print(text)
}

func calculateFutureValues (investmentAmount, expectedReturnRate, years float64) (fV float64, fRV float64) {
	fV = investmentAmount * math.Pow((1 + expectedReturnRate / 100), years)
	fRV = fV / math.Pow(1 + inflationRate / 100, years)
	return
}

func printOutputText (firstText, secondText string) {
	fmt.Println("======================== Result:")
	fmt.Print(firstText, secondText)
}

func printAditionalInfo () {
	fmt.Print("\n")
	fmt.Println("======================== Acronyms:")
	fmt.Println("AFI - Adjusted for inflation")
	fmt.Print("\n")
}