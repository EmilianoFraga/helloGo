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

	testIntegerPrimitiveTypes()

	testFloatPrimitiveTypes()

	testConst()

	testArrays()

	testSlices()

	testMaps()

	testIf()
	
	testSwitch()
}
