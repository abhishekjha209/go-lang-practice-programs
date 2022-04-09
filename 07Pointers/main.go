package main

import "fmt"

func main() {
	fmt.Println("Pointers in nutshell")

	// var genPtr *int;
	// fmt.Println(genPtr); // Contains <nil>

	var value int = 5
	var ptr *int = &value

	// value:var	ptr:address of var	*ptr:derefrencing/value@ptr
	fmt.Println(value, ptr, *ptr)

}
