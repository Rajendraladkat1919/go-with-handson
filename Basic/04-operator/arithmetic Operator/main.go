package main

import "fmt"

func main() {
	a := 10
	b := 20

	var str1 string = "Pune"
	var str2 string = " Mumbai"

	fmt.Println("a + b:", a+b)
	fmt.Println("Addition of strings:", str1+str2) // concatenation
	fmt.Println("a - b:", a-b)
	fmt.Println("a * b:", a*b)
	fmt.Println("b / a:", b/a)
	fmt.Println("b % a:", b%a) // modulus

	var aStr1, bStr2_ = "foo", "bar"
	fmt.Println("concatenation of two strings", aStr1+bStr2_)
}
