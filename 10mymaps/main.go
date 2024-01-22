package main

import "fmt" 

func main(){
	// Maps:- the length is of variable
	// Maps are key-value pairs
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("List of all Languages:- ",languages)
	fmt.Println("JS in short is:- ",languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all Languages:- ",languages)

	// Loops 
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}