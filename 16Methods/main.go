package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	abhishek := User{"abhishek", "abhishek@go.dev", true, 16}
	fmt.Println(abhishek) 
	fmt.Printf("abhishek details are: %+v\n", abhishek)
	fmt.Printf("Name is %v and email is %v.\n", abhishek.Name, abhishek.Email)
	abhishek.GetStatus()
	abhishek.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", abhishek.Name, abhishek.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
