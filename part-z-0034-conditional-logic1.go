package main

import (
	"fmt"
	"strings"
)

func main() {

	var age uint8 = 22

	country := "Canada"

	job := "Doctor"

	if age > 18 && strings.ToLower(country) == "canada" {

		fmt.Println("You can vote!")
	}

	if strings.ToLower(job) == "doctor" || strings.ToLower(job) == "teacher" {

		fmt.Println("You're not a soldier!")
	}

}

// You can vote!
// You're not a soldier!
