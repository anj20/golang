package main

import "fmt"

func main() {
	fmt.Println("Structs")
	// Structs:- the length is of variable
	// No concept of Inheritance in GoLang
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 21}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and Email is %v\n", hitesh.Name, hitesh.Email)
	active := hitesh.GetStatus()
	fmt.Println("The User is Active:- ",active)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int	
} 

func (u User) GetStatus() string{
	fmt.Println("GetStatus function is called")
	if u.Status == true {
		return "Active"
	}
	return "Inactive"
}