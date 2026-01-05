package main

import (
	"fmt"
)

func main() {

	var nums = [4][3]int8{{-5, -2, -99}, {10, 100, 50}, {-22, -33, -44}, {5, 45, 95}}

	for i:=0; i < 4; i++ {

		for j:=0; j < 3; j++ {

			fmt.Printf("Row_%d - number: %d\n", i + 1, nums[i][j])
		}
		fmt.Println("-----------")
	}
}

// Row_1 - number: -5
// Row_1 - number: -2
// Row_1 - number: -99
// -----------
// Row_2 - number: 10
// Row_2 - number: 100
// Row_2 - number: 50
// -----------
// Row_3 - number: -22
// Row_3 - number: -33
// Row_3 - number: -44
// -----------
// Row_4 - number: 5
// Row_4 - number: 45
// Row_4 - number: 95
// -----------
