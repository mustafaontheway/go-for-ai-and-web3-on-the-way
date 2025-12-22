package main

import "fmt"

func main() {

	sum := 0

	var sum_typed uint8 = 0

	for i := 1; i <= 3; i++ {

		sum += i
		fmt.Println("Sum: ", i)
	}

	fmt.Println("..........................")

	for i := uint8(1); i <=10; i += 2 {

		sum_typed += i
		fmt.Println("Sum is: ", sum_typed)
	}
	
}

// Sum:  1
// Sum:  2
// Sum:  3
// ..........................
// Sum is:  1
// Sum is:  4
// Sum is:  9
// Sum is:  16
// Sum is:  25
