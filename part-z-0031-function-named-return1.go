package main

import (
	"fmt"
)

func main() {

	s, d, m := calculate(10, -50)

	fmt.Println(s, d, m) // -40 60 -500
}


func calculate(x int, y int) (sum int, diff int, mult int) {

	sum = x + y 
	diff = x - y
	mult = x * y

	return
}
