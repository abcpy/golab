package main

import "fmt"

func main() {
	sex := 1
	sexStr := ""
	switch sex {
	case 1:
		sexStr = "Male"
	case 2:
		sexStr = "FeMale"
	default:
		sexStr = "Unkown"
	}
	fmt.Println(sexStr)

	score := 80
	grade := ""
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80 && score <= 89:
		grade = "B"
	case score >= 70 && score <= 79:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println(grade)

	
}
