package main

import "fmt"

func main() {
	var number int = 800
	switch number {
	case 90:
		fmt.Println("The number is 90")
	case 80:
		fmt.Println("The number is 80")

	case 50, 100:
		fmt.Println("The number is either 50 or 100")
	case 101:
		fmt.Println("The number is 101")
	default:
		fmt.Println("The number is not in the list")
	}
}
