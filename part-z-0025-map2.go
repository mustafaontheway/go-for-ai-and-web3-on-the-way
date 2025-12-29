package main

import (
	"fmt"
)

func main() {

	cities := map[uint8] string {35: "İzmir", 34: "İstanbul", 6: "Ankaraaa", 46: "Kahramanmaraş"}

	cities[6] = "Ankara" // update a city

	cities[38] = "Kayseri" // add a new city

	delete(cities, 34) // delete a city

	fmt.Println(cities) // map[6:Ankara 35:İzmir 38:Kayseri 46:Kahramanmaraş]
}


