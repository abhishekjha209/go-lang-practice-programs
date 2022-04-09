package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Go")

	// To declare Slices in Go, syntax will be same as array, just don't mention size along w it.
	var fruitList = []string{"Apple", "Banana", "Peaches"}
	fmt.Printf("Type of array %T \n", fruitList)

	// Appending items to the slices.
	fruitList = append(fruitList, "Peach", "Cherry")
	fmt.Println(fruitList)

	// Slicing in slices.
	fmt.Println(fruitList[1:])
	fmt.Println(fruitList[:3])
	fmt.Println(fruitList[1:2])

	highscores := make([]int, 4)

	highscores[0] = 65
	highscores[1] = 45
	highscores[2] = 55
	highscores[3] = 35
	// highscores[4] = 45

	highscores = append(highscores, 66, 77)
	fmt.Println(highscores)

	sort.Ints(highscores)
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))

}
