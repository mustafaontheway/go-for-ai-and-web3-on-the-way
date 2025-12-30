package main

import (
	"fmt"
)

func main() {

	// mem addr is not CONSTANT!

	var year uint16 = 2025

	fmt.Println("Memory address: ", &year) // 0xc00000a098

	var memoryAddress = &year

	fmt.Println(memoryAddress) // 0xc00000a098

	fmt.Println(*memoryAddress) // 2025

	*memoryAddress = 2026

	fmt.Println(year) // 2026
}

