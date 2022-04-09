package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("What rating will you give to our pizza")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our pizza: ")

	// comma ok | comma error syntax.
	// one can consider first part as try another as catch.
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of input is %T, ", input)

}
