package main

import (
	"fmt"
)

func main() {

	sum, diff, mult := calculator(-55, 62)

	fmt.Println(sum, diff, mult) // 7 -117 -3410
}

func calculator(x int64, y int64) (int64, int64, int64) {

	return x + y, x - y, x * y
}
