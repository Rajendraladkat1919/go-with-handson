package main

import "fmt"

func main() {
	number := 2
	switch number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three (because of fallthrough)")
	default:
		fmt.Println("Default")
	}
}
