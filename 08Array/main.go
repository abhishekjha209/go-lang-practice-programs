package main

import "fmt"

func main() {
	fmt.Println("Array in nutshell.")

	// Declaration of array in go requires proper mentioning of size and type of the req array.
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Peach"
	fruitList[3] = "Banana"

	fmt.Println("FruitList is ", fruitList)
	fmt.Println("FruitList is of length ", len(fruitList))

	// Declaring array and initialising at the same time.
	var vegyList = [8]string{"Brinjal", "Beans", "Potato", "Mushroom"}
	fmt.Println(vegyList, len(vegyList))

	fmt.Printf("Type of array %T \n", vegyList)

}

// We don't use array enough in our programs as they don't benifit us in long run.
// It doesn't have cool functions like searching, sorting etc.
