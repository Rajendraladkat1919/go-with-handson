package main

import "fmt"

func main() {
	var b byte = 'A' // byte is an alias for uint8

	fmt.Println("Byte as char:", string(b)) // Output: A
	fmt.Println("Byte value:", b)           // Output: 65 (ASCII)
}
