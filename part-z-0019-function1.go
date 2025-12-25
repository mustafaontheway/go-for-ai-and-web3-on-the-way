package main

import (
	"fmt"
)

func main() {

	greet("Mustafa") // Hello, Mustafa!

	sum(-6500, 45789, 6999) // -6500 + 45789 + 6999 = 46288
}

func greet(name string) {

	fmt.Printf("Hello, %s!\n", name) 
}

func sum(x int64, y int64, z int64) {

	fmt.Printf("%d + %d + %d = %d", x, y, z, x + y + z)
}
