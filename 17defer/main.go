package main

import "fmt"

func main() {
	defer fmt.Println("Three")  // brings to the bottom of the stack
	defer fmt.Println("Two")    
	defer fmt.Println("One")  
	fmt.Println("Hello")
	myDefer()
	fmt.Println()
}

func myDefer(){
	defer fmt.Println()
	for i := 1; i <= 5; i++ {
		defer fmt.Print(i," ")
	}
}