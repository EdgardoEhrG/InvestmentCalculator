package main

import (
	"fmt"
	"math"
)

func printText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fV float64, fRV float64) {
	fV = investmentAmount * math.Pow((1 + expectedReturnRate / 100), years)
	fRV = fV / math.Pow(1 + inflationRate / 100, years)
	return
}

func printOutputText(firstText, secondText string) {
	fmt.Println("======================== Result:")
	fmt.Print(firstText, secondText)
}

func printAditionalInfo() {
	fmt.Print("\n")
	fmt.Println("======================== Acronyms:")
	fmt.Println("AFI - Adjusted for inflation")
	fmt.Print("\n")
}