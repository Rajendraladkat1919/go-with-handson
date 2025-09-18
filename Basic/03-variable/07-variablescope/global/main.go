/*
The scope of the variable is the entire package, and it can be accessed from any function within the same package.
Its scope decide wheather a variable is accessible or not and modifiable or not.
*/
package main

import "fmt"

var globalCounter int = 0

func incrementCounter() {

	fmt.Println("I am from the increment function and Current Counter:", globalCounter)
	println()
	globalCounter++
	fmt.Println(" I am from the increment function and counter after increment:", globalCounter)
}

func main() {

	fmt.Println(" Initial Counter from main:", globalCounter)

	incrementCounter()
	incrementCounter()
	incrementCounter()

	fmt.Println("Final Counter from main:", globalCounter)

}
