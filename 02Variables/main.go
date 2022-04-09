package main

import "fmt"

// One must declare publicaly access variables with first leeter as capital.
const FixedValue = 890

func main() {
	fmt.Println("Variables")

	var userName string = "Abhishek"
	fmt.Println(userName)
	fmt.Printf("Variable is of type %T \n", userName)

	var isLoggedIn bool = true
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var value int = 32879
	fmt.Printf("Variable is of type %T \n", value)

	var smallFloat float32 = 912.34524892
	fmt.Printf("Variable is of type %T \n", smallFloat)

	// float 64 has more precision.
	var bigFloat float64 = 912.34524892
	fmt.Printf("Variable is of type %T \n", bigFloat)

	var anotherVariable int // not initialised. (Remember: Go doesn't put garbage values by itself.)
	fmt.Println(anotherVariable)

	// implicit way of declaring variables
	var website = "github.com/abhishekjha209"
	fmt.Println(website)

	// walrus declaration or no var-keyword declaration
	// walrus declaration can only be used inside a method.
	numberOfUser := 5000
	fmt.Println(numberOfUser)

	fmt.Println(FixedValue)

}
