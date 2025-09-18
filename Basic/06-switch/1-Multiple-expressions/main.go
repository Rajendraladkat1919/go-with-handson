package main

import "fmt"

func main() {
	day := "Saturday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")
	default:
		fmt.Println("Invalid day")
	}
}
