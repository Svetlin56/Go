package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	investmentAmount := 1000.0
	years := 10.0
	expectedReturnRate := 5.3

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
