package main

import (
	"fmt"
)

func main() {

	ages := []uint8 {99, 22, 10, 65, 77}

	for i, age := range ages {

		fmt.Printf("Age %d: %d\n", i + 1, age)
	}

}

// Age 1: 99
// Age 2: 22
// Age 3: 10
// Age 4: 65
// Age 5: 77

