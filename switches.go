package main

import "fmt"


func testSwitch() {
	testSwitchType(5)

	testSwitchType(5.3)

	testSwitchType(true)

	v := [...]int {1, 2 ,3}
	testSwitchType(v)

	s := v[:]
	testSwitchType(s)

	testSwitchCases()
}

func testSwitchType(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Println("Type is int", t)
	case float64:
		fmt.Println("Type is float64", t)
	case bool:
		fmt.Println("Type is boolean", t)
	case [3]int:
		fmt.Println("Type is array [3] of int", t)
	case []int:
		fmt.Println("Type is slice of int", t)
	}
}

func testSwitchCases() {
	v := 10

	switch {
	case v < 5:
		fmt.Println("v is lesser than 5")
	case v < 10:
		fmt.Println("v is lesser than 10")
	case v < 15:
		fmt.Println("v is lesser than 15")
	case v < 20:
		fmt.Println("v is lesser than 20")
	case v < 25:
		fmt.Println("v is lesser than 25")
	default:
		fmt.Println("v is greater than or equal to 25")
	}

	switch v {
	case 1, 2, 3, 4:
		fmt.Println("v is lesser than 5")
	case 5, 6, 7, 8, 9:
		fmt.Println("v is lesser than 10")
	case 10, 11, 12, 13, 14:
		fmt.Println("v is lesser than 15")
	case 15, 16, 17, 18, 19:
		fmt.Println("v is lesser than 20")
	case 20, 21, 22, 23, 24:
		fmt.Println("v is lesser than 25")
	default:
		fmt.Println("v is greater than or equal to 25")
	}
}