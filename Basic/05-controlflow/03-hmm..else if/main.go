package main

import "fmt"

func main() {
	var a, b, c int = 10, 20, 30

	if a > b {
		fmt.Println("a is greater than b")
	} else if b > c {
		fmt.Println("b is greater than c")
	} else if c > a {
		fmt.Println("c is greater than a")
	} else {
		fmt.Println("All are equal")
	}

	fmt.Println("End of the program")
}
