package main

import "fmt"

// Variables examples
func main() {

	// Variables that init and auto detecting types
	var stringVar = "Variable that init and automatic find type of that variable"
	fmt.Println(stringVar)

	intVar := 3
	fmt.Println(intVar)

	var oneOf, twoOf = 1, 2
	fmt.Println(oneOf, twoOf)

	// Variables that you init and definite their type
	var intVarStrict int = 3
	fmt.Println(intVarStrict)

	// Variables that not init
	var intVarStrictNotInit int
	fmt.Println(intVarStrictNotInit)

}
