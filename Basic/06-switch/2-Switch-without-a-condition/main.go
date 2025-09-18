package main

import "fmt"

func main() {
	age := 20
	switch {
	case age < 13:
		fmt.Println("Child")
	case age >= 13 && age < 18:
		fmt.Println("Teenager")
	case age >= 18:
		fmt.Println("Adult")
	}
}
