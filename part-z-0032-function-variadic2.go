package main

import (
	"fmt"
)

func main() {

	m1 := members("Ayhan", "Aykan", "Aybilge", "Aygün")

	fmt.Println("Last member is:", m1[len(m1) - 1]) // Last member is: Aygün
}

func members(m ...string) ([]string) {

	return m
}

