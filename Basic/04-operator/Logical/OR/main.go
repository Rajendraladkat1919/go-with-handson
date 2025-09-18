package main

import "fmt"

func main() {
	var y int = 20
	fmt.Println(y > 5 || y < 15) // true || false = true
	fmt.Println(y > 5 || y > 15) // true || true = true
	fmt.Println(y < 5 || y < 15) // false || false = false
	fmt.Println(y < 5 || y > 15) // false || true = true
}
