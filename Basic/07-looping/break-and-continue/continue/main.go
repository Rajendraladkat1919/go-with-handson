package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Println("Skipping", i)
			continue
		}
		fmt.Println("i =", i)
	}
}
