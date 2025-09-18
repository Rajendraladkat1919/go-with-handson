package main

import (
	"fmt"
)

func localIncrementCounter() {
	localCounter := 0
	println()
	fmt.Println(" Current Counter from local function:", localCounter)
	println()
	localCounter++
	fmt.Println(" Local function counter after increment:", localCounter)
}

func main() {

	localCounter := 0
	fmt.Println(" Initial Counter from main:", localCounter)
	localIncrementCounter()
	localIncrementCounter()

	fmt.Println("Final Counter from main:", localCounter)

} /*
The scope of the variable is limited to the block in which it is defined, such as within a function or a loop.
A variable defined inside a function cannot be accessed from outside that function.
*/
