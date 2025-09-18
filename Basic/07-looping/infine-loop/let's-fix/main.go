package main

// If we dont provide condition in for loop it will be infinite loop
// We can use break statement to exit from infinite loop

import "fmt"

func main() {

	sum := 0

	for {
		sum++

		if sum == 10 {
			break
		}
	}

	fmt.Println("Looking Sexy", sum)

	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("Cool Dude", i)
	}

}
