package main

import "fmt"

func main() {

	var year uint16 = 2013

	for year < 2018 {

		fmt.Println("Year is ", year)

		year++
	}
	
}

// Year is  2013
// Year is  2014
// Year is  2015
// Year is  2016
// Year is  2017
