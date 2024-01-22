package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)
	// for loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for index, day := range days {
		fmt.Printf("Index: %d, Day: %s \n", index+1, day)
	}
}
