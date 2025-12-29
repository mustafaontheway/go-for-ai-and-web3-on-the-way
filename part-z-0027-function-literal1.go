package main

import (
	"fmt"
)

func main() {

	var printCity = func (city string)  {
		
		fmt.Printf("City is %s.\n", city)
	}

	printCity("Ankara")
	printCity("İzmir")
}

// City is Ankara.
// City is İzmir.
