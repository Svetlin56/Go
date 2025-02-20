package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 750
	var expectedReturnRate = 5.3
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
