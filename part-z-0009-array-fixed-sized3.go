package main

import "fmt"

func main() {

	var ages = [...]uint8{12, 9, 99, 41, 21, 77, 96, 81, 40}

	ages[4] = 10

	fmt.Println(ages[len(ages) - 1])

	fmt.Println(ages)
	
}

// 40
// [12 9 99 41 10 77 96 81 40]
