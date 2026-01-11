package main

import (
	"fmt"
) 

func main() {

	var counter uint8 = 10 

	for counter >= 5 {

		fmt.Println("Counter is equal to:", counter)

		counter--
	}
}

// Counter is equal to: 10
// Counter is equal to: 9
// Counter is equal to: 8
// Counter is equal to: 7
// Counter is equal to: 6
// Counter is equal to: 5
