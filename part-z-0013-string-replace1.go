package main

import (
	"fmt"
	"strings"
)

func main() {

	var city1 = "Ankare"

	fmt.Println(strings.Replace(city1, "e", "a", 1)) // Ankara

	city2 := "Ankareeee"

	city2 = strings.Replace(city2, "e", "a", 2)

	fmt.Println(city2) // Ankaraaee
}

