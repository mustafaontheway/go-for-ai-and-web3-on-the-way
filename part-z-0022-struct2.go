package main

import (
	"fmt"
)

func main() {

	empAyhan := Employee{"Ayhan Bilir", "HR", "ab004296", 4_600}

	fmt.Println(empAyhan) // {Ayhan Bilir HR ab004296 4600}
}

type Employee struct {
	name string
	department string
	id string
	salaryUSD uint16
}
