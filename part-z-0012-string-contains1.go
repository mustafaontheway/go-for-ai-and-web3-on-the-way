package main

import (
	"fmt"
	"strings"
)

func main() {

	var city = "Ankara"

	var color = "kara"

	fmt.Println(strings.Contains(city, color)) 

	fmt.Println(strings.Contains(color, city)) 
}

// true
// false
