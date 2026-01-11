package main

import (
	"fmt"
) 

func main() {

	fmt.Println(friends(220, 284)) // true

	fmt.Println(friends(464, 962)) // false
}

func friends(x uint, y uint) bool {

	var sum1 uint
	var counter1 uint
	var sum2 uint 
	var counter2 uint

	for counter1 = 1; counter1 < x; counter1++ {

		if x % counter1 == 0 {

			sum1 += counter1
		}
	}

	for counter2 = 1; counter2 < y; counter2++ {

		if y % counter2 == 0 {

			sum2 += counter2
		}
	}

	if x == sum2 && y == sum1 {

		return true
	}

	return false
}
