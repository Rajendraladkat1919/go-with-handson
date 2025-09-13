package main

import "fmt"

func main() {

	var ua uint8 = 254
	var ub uint16 = 32767
	var uc uint32 = 214748
	var ud uint64 = 9223372036854

	fmt.Println("Unsigned Integers:", ua) // Output: Unsigned Integers: 138
	fmt.Println("Unsigned Integers:", ub) // Output: Unsigned Integers: 32768
	fmt.Println("Unsigned Integers:", uc) // Output: Unsigned Integers: 2147483648
	fmt.Println("Unsigned Integers:", ud) // Output: Unsigned Integers: 9223372036854775808
}
