// compare two values and return boolean result
// allow values of same type only
// ++ != < > <= >=

package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println("a < b:", a < b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a <= b:", a <= b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a == b:", a == b)

	var city string = "Pune"
	var city1 string = "Mumbai"
	fmt.Println("city != city1:", city != city1)
	fmt.Println("city < city1:", city < city1)   // lexicographical comparison
	fmt.Println("city > city1:", city > city1)   // lexicographical comparison
	fmt.Println("city <= city1:", city <= city1) // lexicographical comparison
	fmt.Println("city >= city1:", city >= city1) // lexicographical comparison
	fmt.Println("city == city1:", city == city1) // lexicographical comparison

}
