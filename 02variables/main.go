package main

import "fmt"

// jwtToken := "asdasdasd"
// not allowed outside of function
const LoginToken string = "asdasdasd" // Public


func main() {
	fmt.Println("Variables")
	var username string = "root"
	fmt.Println(username)
	fmt.Printf("Variable username is of type: %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable isLoggedIn is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable smallVal is of type: %T\n", smallVal)

	// default values and some aliases
	var anothervar string
	fmt.Println(anothervar)
	fmt.Printf("Variable anothervar is of type: %T\n", anothervar)


	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("Variable website is of type: %T\n", website)

	// no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("The type of LoginToken is: %T\n", LoginToken)
}
