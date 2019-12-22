package main

import "fmt"

func testForOne() {
	i := 500

	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	for i := 9; i >= 0; i-- {
		fmt.Println("i = ", i)
	}

	fmt.Println("outer i = ", i)

	for i > 0 {
		fmt.Println("i = ", i)

		i /= 2
	}

	fmt.Println("outer i = ", i)

	slice := []int {12, 3, 49, 21, 44, 45}
	for i, x := range slice {
		if i > 0 { print(", ")}
		fmt.Print("x = ", x)
	}
	fmt.Println()
}