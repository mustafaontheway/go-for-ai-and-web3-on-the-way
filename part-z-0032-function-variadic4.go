package main

import (
	"fmt"
) 

func main() {

	var x, y, z uint8 = 33, 24, 48

	fmt.Println(findMaxNum(x, y, z)) // 48

}

func findMaxNum(nums ...uint8) uint8 {

	var maxNum uint8

	for _, num := range nums {

		if num > maxNum {

			maxNum = num
		}
	}

	return maxNum
}
