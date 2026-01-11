package main

import (
	"fmt"
) 

func main() {

	name := "Mustafa"

	for index := 0; index < len(name); index++ {

		fmt.Printf("-%c-", name[index])
	}
}

// -M--u--s--t--a--f--a-
