package main

import (
	"fmt"
)

func main() {

	// slices are dynamic...

	years := [6]uint16{2000, 2005, 2007, 2009, 1978, 1654}

	var names = [5]string{"ayhan", "hakan", "kağan", "aygün", "bengü"}

	yearsSlice1 := years[:]

	yearsSlice2 := years[1:4]
	
	namesSlice := names[2:]

	fmt.Println(yearsSlice1)
	fmt.Println(yearsSlice2)
	fmt.Println(namesSlice)
}

// [2000 2005 2007 2009 1978 1654]
// [2005 2007 2009]
// [kağan aygün bengü]
