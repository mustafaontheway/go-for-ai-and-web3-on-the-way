package main

import (
	"fmt"
)

func main() {

	var sales [4][3]uint32

	sales[0][0] = 55000
	sales[0][1] = 49000
	sales[0][2] = 67000

	fmt.Println(sales) // [[55000 49000 67000] [0 0 0] [0 0 0] [0 0 0]]

}

