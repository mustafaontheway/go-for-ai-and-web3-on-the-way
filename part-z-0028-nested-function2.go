package main

import (
	"fmt"
)

func main() {

	printCity("İzmir")

}

func printCity(city string) {

	var greet = func ()  {
		
		fmt.Printf("Welcome to %s!\n", city)
	}

	greet()

	fmt.Printf("%s is a nice city.\n", city)
}

// Welcome to İzmir!
// İzmir is a nice city.
