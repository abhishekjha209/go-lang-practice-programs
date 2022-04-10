package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	// Adding items in map
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	// Accessing items in map
	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	// Deleting items in map (using key.)
	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
