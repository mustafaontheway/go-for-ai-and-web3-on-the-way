package main

import (
	"fmt"
)

func main() {

	cities := map[uint8] string {35: "İzmir", 34: "İstanbul", 6: "Ankara", 46: "Kahramanmaraş"}

	for key, city := range cities {

		fmt.Printf("City %d is %s\n", key, city)
	}
}

// City 35 is İzmir
// City 34 is İstanbul
// City 6 is Ankara
// City 46 is Kahramanmaraş
