// Please note that multiple variables deceralation and assignment done with = operator not with := operator.

package main

import "fmt"

func main() {
	var age, code int = 25, 1001
	var height, weight, BMI float32 = 5.9, 76.30, 22.5
	var name, city string = "John Doe", "New York"
	var isStudent, isEmployed bool = false, true

	fmt.Println("Name:", name, "City:", city)
	fmt.Println("Age:", age, "Code:", code)
	fmt.Println("Height:", height, "Weight:", weight, "BMI:", BMI)
	fmt.Println("Is Student:", isStudent, "Is Employed:", isEmployed)
}
