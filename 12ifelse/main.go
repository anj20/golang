package main

import (
	"fmt"
)

func main() {
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login counts"
	}

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}

	if num := 3; num < 10 {	
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}


	fmt.Println(result)
}
