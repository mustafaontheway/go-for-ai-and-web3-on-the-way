package main

import (
	"fmt"
)

func main() {

	years := make([]uint16, 10)

	fmt.Println(years) // [0 0 0 0 0 0 0 0 0 0]

	years[4] = 2000

	fmt.Println(years) // [0 0 0 0 2000 0 0 0 0 0]

	//years[11] = 2050 // panic: runtime error: index out of range [11] with length 10

	years = append(years, 2050)

	fmt.Println(years) // [0 0 0 0 2000 0 0 0 0 0 2050]

	fmt.Println(cap(years)) // 24

	fmt.Println(len(years)) // 11
}

