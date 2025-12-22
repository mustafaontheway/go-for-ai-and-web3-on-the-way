package main

import "fmt"

func main() {

	var year uint16 = 2013

	for year < 2025 {

		if year % 2 == 0 {

			fmt.Println("Even year: ", year)

		} else {

			fmt.Println("Odd year: ", year)
		}

		year++
	}
	
}

// Odd year:  2013
// Even year:  2014
// Odd year:  2015
// Even year:  2016
// Odd year:  2017
// Even year:  2018
// Odd year:  2019
// Even year:  2020
// Odd year:  2021
// Even year:  2022
// Odd year:  2023
// Even year:  2024
