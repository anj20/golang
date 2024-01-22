package main

import "fmt"

func main() {
	fmt.Println("Array")

	var arr [5]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println("The Array list is: ",  arr)
	fmt.Println("The Array list is: ", len(arr))

	var vegList = [5]string{"Potato", "Tomato", "Onion", "Carrot", "Cabbage"}	
	fmt.Println("The Array list is: ",  vegList) 
}
