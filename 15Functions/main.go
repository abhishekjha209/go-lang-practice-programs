package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 5, 8, 7, 3)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro Message is: ", myMessage)

}

// specifying return type as well.
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// Example of Variadic functions:
// It can be called with any number of trailing arguments. 
// For example, fmt.Println is a common variadic function.
func proAdder(values ...int) (int, string) {

	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result function"
}

func greeter() {
	fmt.Println("Namstey from golang")
}
