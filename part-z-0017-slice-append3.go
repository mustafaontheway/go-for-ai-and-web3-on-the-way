package main

import (
	"fmt"
)

func main() {

	years := []uint16{2000, 2005, 2007, 2009, 1978, 1654, 2001}

	var evenYears []uint16
	var oddYears []uint16

	// evenYears := []uint16{}
	// oddYears := []uint16{}

	for i := 0; i < len(years); i++ {

		if years[i] % 2 == 0 {

			evenYears = append(evenYears, years[i])
		} else {

			oddYears = append(oddYears, years[i])
		}
	}

	fmt.Println("Even years:", evenYears)
	fmt.Println("Odd years:", oddYears)
}

// Even years: [2000 1978 1654]
// Odd years: [2005 2007 2009 2001]
