package main

import "fmt"

func main() {
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				fmt.Println("Continuing outer loop at j =", j)
				continue outer
			}
			fmt.Println("i =", i, "j =", j)
		}
	}
}
