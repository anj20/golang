package main

import "fmt"

func main(){
	fmt.Println("Hello World")
	greeter()
	// func greetertwo(){
	// 	fmt.Println("Namaste for golang2")
	// }
	// not possible
	ans := adder(2,3)
	ans2 := adder2(2,3,4,5,6,7,8,9)
	fmt.Println("Normal Sum is:-",ans)
	fmt.Println("Pro result is:- ",ans2) 
}

func adder(value1 int, value2 int) int{
	return value1 + value2
}
func adder2(values ...int ) int{
	total := 0
	for _, value := range values{
		total += value
	}
	return total
}
func greeter(){
	fmt.Println("Namaste for golang")
}