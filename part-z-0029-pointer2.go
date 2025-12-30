package main

import (
	"fmt"
)

func main() {

	// mem addr is not CONSTANT!

	var year uint16 = 2025

	var pointerYear *uint16

	pointerYear = &year

	fmt.Println(pointerYear) // 0xc00000a098

	fmt.Println(*pointerYear == year) // true

	*pointerYear = 2026

	fmt.Println(year) // 2026
}

