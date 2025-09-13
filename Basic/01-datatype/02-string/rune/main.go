package main

import "fmt"

func main() {
	var ch rune = 'G' // Single character in single quotes

	fmt.Println("Character:", string(ch))  // Convert rune to string for display
	fmt.Println("Unicode Code Point:", ch) // Shows Unicode value (71)
}
