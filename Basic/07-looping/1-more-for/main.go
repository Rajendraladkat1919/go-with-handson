package main

import "fmt"

func main() {
	var i int
	i = 1

	// i :=0
	for i <= 10 {
		fmt.Println(i * i)
		i += 2
	}
}
