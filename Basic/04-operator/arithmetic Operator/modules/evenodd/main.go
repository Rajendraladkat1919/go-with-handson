package main

import "fmt"

func main() {

	var number int = 20
	if number%2 == 0 {
		fmt.Println(number, "is even number")
	} else {
		fmt.Println(number, "is odd number")
	}
}
