package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.4

	//fmt.Print("Enter investment amount: ")
	outputText("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Enter expected return rate: ")
	outputText("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Enter years: ")
	outputText("Enter years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for Inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	//return
}
