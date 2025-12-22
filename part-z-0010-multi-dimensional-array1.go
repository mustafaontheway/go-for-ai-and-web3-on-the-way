package main

import "fmt"

func main() {

	var m = [3][3]int8{{-4, 11, 22}, {12, -6, 21}, {45, -54, 100}}

	fmt.Println(m) 

	fmt.Println("Max number: ", m[2][2])

	fmt.Println("Min number: ", m[2][1])
}

// [[-4 11 22] [12 -6 21] [45 -54 100]]
// Max number:  100
// Min number:  -54
