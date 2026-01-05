package main

import (
	"fmt"
)

func main() {

	var salesJanuary uint32 = 650_957

	var premiumJanuary = calculatePremium(salesJanuary)

	fmt.Printf("January premium %.2f ₺\n", premiumJanuary)

	var salesFebruary uint32 = 387_654

	fmt.Printf("February premium %.2f\n", calculatePremium((salesFebruary)))
}

func calculatePremium(sales uint32) float32 {

	if sales > 500_000 {

		return float32(sales) * 0.03
	} 

	return 0.0
}

// January premium 19528.71 ₺
// February premium 0.00
