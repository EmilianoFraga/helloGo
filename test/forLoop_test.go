package main_test

import (
	"fmt"
	"testing"
)

func TestForOne(t *testing.T) {
	const expectedI = 500

	i := expectedI

	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	if i != expectedI {
		t.Errorf("Invalid %d expected as %d", i, expectedI)
	}

	for i := 9; i >= 0; i-- {
		fmt.Println("i = ", i)
	}

	if i != expectedI {
		t.Errorf("Invalid %d expected as %d", i, expectedI)
	}

	for i > 0 {
		fmt.Println("i = ", i)

		i /= 2
	}

	if i != 0 {
		t.Errorf("Invalid %d expected as %d", i, 0)
	}

	slice := []int {12, 3, 49, 21, 44, 45}
	for i, x := range slice {
		if i > 0 { print(", ")}
		fmt.Print("x = ", x)
	}
	fmt.Println()

	if i != 0 {
		t.Errorf("Invalid %d expected as %d", i, 0)
	}
}