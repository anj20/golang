package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of pointers")
	// START OMIT 
	var p *int
	fmt.Println("Value of Pointer p is", p)
	mynumber := 5
	p = &mynumber
	fmt.Println("Value of Pointer p is", p)
	fmt.Println("Value of Pointer p is", *p)
	// END OMIT
}
