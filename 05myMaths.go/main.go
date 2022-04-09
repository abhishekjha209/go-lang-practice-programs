package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in GoLang")

	var myNumber int = 5
	var myfloat float64 = 5.8
	fmt.Println(myNumber + int(myfloat)) // 10 (converting into int, makes us loose accuracy)

	// // Generating Random number using Math:Rand()

	// // seed to get different result every time
	// rand.Seed(time.Now().UnixNano())
	// sample := rand.Intn(5)
	// fmt.Println("Random number using Math:Rand() "sample)

	// Generating Random number using Crypto:Rand()
	sample2, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("Random number using Crypto:Rand()", sample2)
}
