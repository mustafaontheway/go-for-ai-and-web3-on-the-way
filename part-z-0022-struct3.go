package main

import (
	"fmt"
)

func main() {

	empAyhan := Employee{"Ayhan Bilir", "HR", "ab004296", 4_600}

	fmt.Printf("His name is %s.", empAyhan.name) // His name is Ayhan Bilir.
}

type Employee struct {
	name string
	department string
	id string
	salaryUSD uint16
}
