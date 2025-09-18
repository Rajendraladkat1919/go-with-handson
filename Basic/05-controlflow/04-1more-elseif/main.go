package main

import "fmt"

func main() {
	var fruit string = "Banana"

	if fruit == "apple" {
		fmt.Println("I love the apple.")
	} else if fruit == "Oranges" {
		fmt.Println(" Its not the apple, its orange.")
	} else if fruit == "Banana" {
		fmt.Println("Its not the apple, its banana.")
	} else {
		fmt.Println("I dont like any fruit.")
	}

}
