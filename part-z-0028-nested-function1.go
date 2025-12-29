package main

import (
	"fmt"
)

func main() {

	printCity("İzmir")

}

func printCity(city string) {

	var greet = func ()  {
		
		fmt.Println("Welcome to our city!")
	}

	greet()

	fmt.Printf("%s is a nice city.\n", city)
}

// Welcome to our city!
// İzmir is a nice city.
