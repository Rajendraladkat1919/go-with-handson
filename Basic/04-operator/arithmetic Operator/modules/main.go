package main

import "fmt"

func main() {

	var a, b, c int = 10, 20, 30

	fmt.Println(a % b)
	fmt.Println(b % c)
	fmt.Println(c % a)
}
