package main

import (
	"fmt"
)

func main() {

	students := make(map[uint16] string)

	students[438] = "Jale Cüce"

	students[429] = "Aydın Uçar"

	fmt.Println(students) // map[438:Jale Cüce 429:Aydın Uçar]

	student_aydin := students[429]

	fmt.Println(student_aydin) // Aydın Uçar
}

