package main

import "fmt"

func main() {
	dict := make(map[string][]string)

	dict["Fruits"] = []string{"Apple", "Banana", "Cherry"}
	dict["Vegitables"] = []string{"curryLeaves", "Brinjal"}

	fmt.Println(dict)

}
