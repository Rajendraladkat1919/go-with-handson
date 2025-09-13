package main

import (
	"fmt"
	"math/cmplx" // Package for complex number operations
)

func main() {
	// Boolean type
	var isTrue bool = true
	var isFalse bool = false
	fmt.Printf("Boolean types: isTrue = %t, isFalse = %t\n", isTrue, isFalse)

	// Integer types
	var anInt int = 42
	var anInt8 int8 = -12
	var aUint uint = 100
	var aUint16 uint16 = 65535
	fmt.Printf("Integer types: anInt = %d, anInt8 = %d, aUint = %d, aUint16 = %d\n", anInt, anInt8, aUint, aUint16)

	// Floating-point types
	var aFloat32 float32 = 3.14
	var aFloat64 float64 = 2.71828
	fmt.Printf("Floating-point types: aFloat32 = %f, aFloat64 = %f\n", aFloat32, aFloat64)

	// Complex types
	var aComplex64 complex64 = 1 + 2i
	var aComplex128 complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Printf("Complex types: aComplex64 = %v, aComplex128 = %v\n", aComplex64, aComplex128)

	// String type
	var aString string = "Hello, Go!"
	fmt.Printf("String type: aString = \"%s\"\n", aString)

	// Rune type (alias for int32, represents a Unicode code point)
	var aRune rune = 'G'
	fmt.Printf("Rune type: aRune = '%c' (Unicode value: %d)\n", aRune, aRune)

	// Byte type (alias for uint8)
	var aByte byte = 255
	fmt.Printf("Byte type: aByte = %d\n", aByte)

	// Array type
	var anArray [3]int = [3]int{1, 2, 3}
	fmt.Printf("Array type: anArray = %v\n", anArray)

	// Slice type
	aSlic := []string{"apple", "banana", "cherry"}
	fmt.Printf("Slice type: aSlic = %v\n", aSlic)

	// Struct type
	type Person struct {
		Name string
		Age  int
	}
	aPerson := Person{"Alice", 30}
	fmt.Printf("Struct type: aPerson = %v\n", aPerson)

	// Pointer type
	var ptr *int
	val := 10
	ptr = &val
	fmt.Printf("Pointer type: ptr = %v, value pointed to by ptr = %d\n", ptr, *ptr)
}
