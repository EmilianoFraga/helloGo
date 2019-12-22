package main

import "fmt"

func testIf() {
	value := 5

	if value < 5 {
		fmt.Println("Value is lesser than 5")
	} else if value > 5 {
		fmt.Println("Value is greater than 5")
	} else {
		fmt.Println("Value is just 5")
	}

	if value := 6; value < 5 {
		fmt.Println("Inner value is lesser than 5")
	} else if value > 5 {
		fmt.Println("Inner value is greater than 5")
	} else {
		fmt.Println("Inner value is just 5")
	}

	fmt.Println("Outer value: ", value)
}