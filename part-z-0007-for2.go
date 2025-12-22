package main

import "fmt"

func main() {

	var sum int

	for i := -10; i <= 20; i++ {

		if i % 2 != 0 {

			sum += i

			fmt.Println("Sum: ", sum)
		}
	}
	
}

// Sum:  -9
// Sum:  -16
// Sum:  -21
// Sum:  -24
// Sum:  -25
// Sum:  -24
// Sum:  -21
// Sum:  -16
// Sum:  -9
// Sum:  0
// Sum:  11
// Sum:  24
// Sum:  39
// Sum:  56
// Sum:  75
