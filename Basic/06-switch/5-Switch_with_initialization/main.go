package main

import "fmt"

func main() {
	switch x := 10; {
	case x < 0:
		fmt.Println("Negative")
	case x == 0:
		fmt.Println("Zero")
	case x > 0:
		fmt.Println("Positive")
	}
}
