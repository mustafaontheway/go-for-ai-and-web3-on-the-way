package main

import (
	"fmt"
	"errors"
)

func main() {

	var year uint16 = 2025

	yearError := errors.New("Initial year must be 2026!")

	if year != 2026 {

		fmt.Println(yearError)
	}
}

// Initial year must be 2026!
