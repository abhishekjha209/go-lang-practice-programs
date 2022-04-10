package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	abhishek := User{"Abhishek", "abhishek@go.dev", true, 16} //initialisation of obj

	// way of printing 1 (only details) :: {Abhishek abhishek@go.dev true 16}
	fmt.Println(abhishek) 

	// way of printing 2 (details as key-value) ::  {Name:Abhishek Email:abhishek@go.dev Status:true Age:16}
	fmt.Printf("hitesh details are: %+v\n", abhishek) 

	// way of printing 3 (as per want)
	fmt.Printf("Name is %v and email is %v.", abhishek.Name, abhishek.Email) 

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
