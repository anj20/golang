package main

import "fmt"

func main() {
	fmt.Println("Structs")
	// Structs:- the length is of variable
	// No concept of Inheritance in GoLang
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 21}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details are: %+v\n", hitesh)
	fm t.Printf("Name is %v and Email is %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
