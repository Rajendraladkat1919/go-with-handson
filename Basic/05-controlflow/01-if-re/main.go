// what if
package main

import "fmt"

func main() {

	var a string = "Pune"

	if a == "Pune" {
		fmt.Println("a is Pune")
	}

	if a == "pune" {
		fmt.Println("I am here")
		fmt.Println("Its not executed.")
	}
	fmt.Println("End of the program")
}
