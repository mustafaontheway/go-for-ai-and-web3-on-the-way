package main

import (
	"fmt"
) 

func main() {

	var a, b, c, d int = -33, 24, -48, 14

	var minNum int = a

	if b < minNum {

		minNum = b
	}

	if c < minNum {

		minNum = c
	}

	if d < minNum {

		minNum = d
	}

	fmt.Println("Min number is:", minNum) // Min number is: -48
}


