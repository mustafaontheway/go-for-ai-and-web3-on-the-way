package main

import (
	"fmt"
)

func main() {

	years := []uint16{2000, 2005, 2007, 2009, 1978, 1654}

	years2 := []uint16{2050, 2051, 2658}

	years = append(years, years2...)

	fmt.Println(years)

}

// [2000 2005 2007 2009 1978 1654 2050 2051 2658]
