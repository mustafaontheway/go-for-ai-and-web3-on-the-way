package main

import (
	"fmt"
)

func main() {

	// slices are dynamic...

	years := []uint16{2000, 2005, 2007, 2009}

	var names = []string{"ayhan", "hakan", "kağan"}

	fmt.Println(years, names) // [2000 2005 2007 2009] [ayhan hakan kağan]
}

