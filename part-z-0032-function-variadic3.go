package main

import (
	"fmt"
)

func main() {

	fmt.Println(mult(1, 3, 5))

	fmt.Println(mult(11, -3, 5, 12))
}

func mult(nums ...int) int {

	result := 1

	for _, num := range nums {

		result *= num
	}

	return result
}

// 15
// -1980
