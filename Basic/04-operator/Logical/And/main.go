package main

import "fmt"

func main() {
	var x int = 10

	fmt.Println(x > 5 && x < 15) // true && true = true
	fmt.Println(x > 5 && x > 15) // true && false = false
	fmt.Println(x < 5 && x < 15) // false && true = false
	fmt.Println(x < 5 && x > 15) // false && false = false
}
