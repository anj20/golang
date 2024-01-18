package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Pizza Shop!")
	fmt.Println("Please rate our pizza between 1 and 5: ")
	reader,_ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Thanks for rating, ", reader) 
	numrating,err := strconv.ParseFloat(strings.TrimSpace(reader), 64)
	if err != nil {
		fmt.Println(err)
		panic("Please enter a number")
	} else {
		fmt.Println("Added 1 to your rating: ", numrating+1)
	}
}
 