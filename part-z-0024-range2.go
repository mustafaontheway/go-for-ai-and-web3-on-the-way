package main

import (
	"fmt"
)

func main() {

	myName := "Mustafa"

	for i, char := range myName {

		fmt.Printf("Char %d: %c\n", i + 1, char)
	}

}

// Char 1: M
// Char 2: u
// Char 3: s
// Char 4: t
// Char 5: a
// Char 6: f
// Char 7: a
