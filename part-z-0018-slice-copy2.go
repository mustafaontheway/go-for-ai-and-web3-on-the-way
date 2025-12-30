package main

import (
	"fmt"
)

func main() {

	years := []uint16{2000, 2008, 2013, 2016, 2020, 2026}

	someYears :=[]uint16{1993, 1997, 2018, 2025}

	copy(someYears, years[2:])

	fmt.Println(someYears) // [2013 2016 2020 2026]

	fmt.Println(years) // [2000 2008 2013 2016 2020 2026]

	copy(someYears, years[:2])

	fmt.Println(someYears) // [2000 2008 2020 2026]
}

