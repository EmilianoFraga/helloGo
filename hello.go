package main

import (
	"fmt"
)

func main() {
	i := 345
	fmt.Println("Hello Go")
	s := fmt.Sprintf("This is a formatted output %05d.", i)
	fmt.Println(s)
	fmt.Println(varinfo(i))
	fmt.Println(varinfo(s))
	fmt.Println(varinfo(s[2]))

	x, y := multiReturn(5)
	fmt.Printf("Multi return: %d %d\n", x, y )

	x, _ = multiReturn(9)
	fmt.Printf("Multi return: %d %d\n", x, y )

	_, y = multiReturn(20)
	fmt.Printf("Multi return: %d %d\n", x, y )

	testIntegerPrimitiveTypes()

	testFloatPrimitiveTypes()

	testConst()

	testArrays()

	testSlices()

	testMaps()

	testIf()

	testSwitch()

	testForOne()
}
