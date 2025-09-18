package main

import "fmt"

func main() {

	var num int = 20
	var num1 int = 10

	fmt.Println("num < num1:", num < num1)
	fmt.Println("num > num1:", num > num1)
	fmt.Println("num <= num1:", num <= num1)
	fmt.Println("num >= num1:", num >= num1)
	fmt.Println("num == num1:", num == num1)
	fmt.Println("num != num1:", num != num1)

	var city string = "Pune"
	var city1 string = "Mumbai"
	fmt.Println("city != city1:", city != city1)
	fmt.Println("city < city1:", city < city1)
	fmt.Println("city == city1:", city == city1)
}
