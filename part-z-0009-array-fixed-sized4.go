package main

import "fmt"

func main() {

	var ages = [...]uint8{12, 9, 99, 41, 21, 77, 96, 81, 40}

	for i := 0; i < len(ages); i++ {

		fmt.Println("Age is ", ages[i])
	}
	
}

// Age is  12
// Age is  9
// Age is  99
// Age is  41
// Age is  21
// Age is  77
// Age is  96
// Age is  81
// Age is  40
