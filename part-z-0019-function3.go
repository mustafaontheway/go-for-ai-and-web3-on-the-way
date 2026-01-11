package main

import (
	"fmt"
) 

func main() {

	fmt.Println(isPrime(16)) // false

	fmt.Println(isPrime(17)) // true

	fmt.Println(isPrime(0)) // false

	fmt.Println(isPrime(1)) // false
}

func isPrime(num uint) bool {

	prime := true

	var counter uint

	if num < 2 {

		prime = false
	}

	for counter = 2; counter < num; counter++ {

		if num % counter == 0 {

			prime = false

			break
		} 
	}

	return prime
}
