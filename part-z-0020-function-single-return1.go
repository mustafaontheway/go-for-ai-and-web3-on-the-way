package main

import (
	"fmt"
)

func main() {

	g1 := greet("Mustafa")

	fmt.Println(g1) // Hello, Mustafa!

	g2 := sum(-88777, 8877, 9966)

	fmt.Println(g2) // -69934
}

func greet(name string) string {

	return "Hello, " + name + "!"
}

func sum(x int64, y int64, z int64) int64 {

	return x + y + z
}
