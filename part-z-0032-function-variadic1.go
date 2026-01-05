package main

import (
	"fmt"
)

func main() {

	members("Ayhan", "Bengü", "Hakan")

	members("Ayhan", "Bengü", "Hakan", "Sevda", "Kağan")
}

func members(m ...string) {

	fmt.Println("Members are:", m)
}

// Members are: [Ayhan Bengü Hakan]
// Members are: [Ayhan Bengü Hakan Sevda Kağan]
