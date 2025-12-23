package main

import (
	"fmt"
	"strings"
)

func main() {

	var fullName = "Mustafa Büyükdereli"

	var nameAndLastName = strings.Split(fullName, " ")

	fmt.Println(nameAndLastName) 

	fmt.Println(nameAndLastName[0])

	fmt.Println(nameAndLastName[1])
}

// [Mustafa Büyükdereli]
// Mustafa
// Büyükdereli
