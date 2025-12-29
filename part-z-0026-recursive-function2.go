package main

import (
	"fmt"
)

func main() {

	result := sumNumbers(50)

	fmt.Println(result)

	result2 := sumNumbers(5000)

	fmt.Println(result2)
}

func sumNumbers(num uint) uint {

	if num < 100 {

		return num + sumNumbers(num - 1)
	}

	return 0
}

// 1275
// 0
