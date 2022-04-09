package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app!")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')	// This will include a endline char as well at last.

	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to the rating: ", numRating+1)
	}

}
