package main

import (
	"fmt"
)

func main() {

	cities := map[uint8] string {35: "İzmir", 34: "İstanbul", 6: "Ankara", 46: "Kahramanmaraş"}
	
	fmt.Printf("The capital of Turkey is %s.", cities[6]) // The capital of Turkey is Ankara.
}


