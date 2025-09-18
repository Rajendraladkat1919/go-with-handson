package main

import "fmt"

func main() {
	// Declaring variables without initialization
	var (
		intValue       int
		floatValue     float64
		boolValue      bool
		stringValue    string
		pointerValue   *int
		sliceValue     []int
		mapValue       map[string]int
		channelValue   chan int
		interfaceValue interface{}
		arrayValue     [3]int
	)

	fmt.Println("Default int:", intValue)
	fmt.Println("Default float64:", floatValue)
	fmt.Println("Default bool:", boolValue)
	fmt.Println("Default string:", stringValue)
	fmt.Println("Default *int (pointer):", pointerValue)
	fmt.Println("Default []int (slice):", sliceValue)
	fmt.Println("Default map[string]int (map):", mapValue)
	fmt.Println("Default chan int (channel):", channelValue)
	fmt.Println("Default interface{}:", interfaceValue)
	fmt.Println("Default [3]int (array):", arrayValue)
}
