package main

import "fmt"

func main() {

	// Declaration without initialization
	var age float32
	var name string
	var isMale bool
	var height int
	// Default values are assigned
	fmt.Println("This is the default value of variables: ", "Name:", name, " Age:", age, " isMale:", isMale, " Height:", height)

	// Assigning values later
	age = 35
	name = "Rala"
	isMale = true
	height = 6
	fmt.Println() // New line for better readability
	fmt.Println("This is the values after assignment", "Name:", name, " Age:", age, " isMale:", isMale, " Height:", height)
}
