package main

import "fmt"

func main() {
	var fruit string = "Papaya"

	if fruit == "Mango" {
		fmt.Println("The Given fruit is not Papaya. Its Mango.")
	} else {
		fmt.Println("The Given fruit is Papaya.")
	}

	fmt.Println("End of the program")
}
