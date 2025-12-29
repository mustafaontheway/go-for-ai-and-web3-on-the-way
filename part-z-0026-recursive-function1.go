package main

import (
	"fmt"
)

func main() {

	printNumbers(6)
}

func printNumbers(num int8) {

	if num > -5 {

		fmt.Println("Number is: ", num)

		printNumbers(num - 1)
	}
}

// Number is:  6
// Number is:  5
// Number is:  4
// Number is:  3
// Number is:  2
// Number is:  1
// Number is:  0
// Number is:  -1
// Number is:  -2
// Number is:  -3
// Number is:  -4
