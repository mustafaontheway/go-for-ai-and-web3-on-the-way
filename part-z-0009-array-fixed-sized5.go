package main

import (
	"fmt"
)

func main() {

	var cities [10]string
	var years [5]uint16

	cities[0] = "Ankara"
	cities[1] = "İzmir"

	years[2] = 2018

	fmt.Println(cities) // [Ankara İzmir        ]

	fmt.Println(years) // [0 0 2018 0 0]

	cities[7] = "Bursa"

	fmt.Println(cities) // [Ankara İzmir      Bursa  ]

	fmt.Println(len(cities)) // 10

}

