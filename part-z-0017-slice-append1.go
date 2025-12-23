package main

import (
	"fmt"
)

func main() {

	// slices are dynamic...

	years := []uint16{2000, 2005, 2007, 2009, 1978, 1654}

	var names = []string{"ayhan", "hakan", "kağan", "aygün", "bengü"}

	years = append(years, 1400, 1440, 1480)

	names = append(names, "mustafa")

	fmt.Println(years)
	fmt.Println(names)
}

// [2000 2005 2007 2009 1978 1654 1400 1440 1480]
// [ayhan hakan kağan aygün bengü mustafa]
