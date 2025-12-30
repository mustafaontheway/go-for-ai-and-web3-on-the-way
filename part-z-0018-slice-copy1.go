package main

import (
	"fmt"
)

func main() {

	years := []uint16{2000, 2008, 2013, 2016, 2020, 2026}

	someYears :=[]uint16{1993, 1997, 2018}

	copy(someYears, years)

	years[0] = 1995

	fmt.Println(someYears) // [2000 2008 2013]

	fmt.Println(years) // [1995 2008 2013 2016 2020 2026]
}

