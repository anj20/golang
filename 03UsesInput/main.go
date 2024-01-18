package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go!"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	// comma ok || err idiom
	input, err := reader.ReadString('\n')
	fmt.Println("Hello, ", input, "nice to meet you!")
	fmt.Printf("Type of this string is %T\n", input)
	fmt.Println("Error: ", err)
}
