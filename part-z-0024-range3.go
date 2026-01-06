package main

import (
	"fmt"
)

func main() {

	years := make([]uint16, 10)

	//fmt.Println(years) // [0 0 0 0 0 0 0 0 0 0]

	years[4] = 2000

	for _, year := range years {

		if year != 0 {

			fmt.Println("Year:", year) // Year: 2000
		}
	}
}

