package main

import "fmt"

func main() {

	var num uint8 

	for i := uint8(1); i < 10; i++ {

		if i == 4 {
			continue
		}

		if i == 8 {
			break
		}

		num += i 

		fmt.Println(i, num)
	}
	
}

// 1 1
// 2 3
// 3 6
// 5 11
// 6 17
// 7 24
