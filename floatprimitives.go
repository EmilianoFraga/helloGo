package main

import (
	"fmt"
)

// testFloatPrimitiveTypes just does... test float primitive types
func testFloatPrimitiveTypes() {
	var minFloat32 float32 = -3.4e+38
	var maxFloat32 float32 = +3.4e+38

	var minFloat64 float64 = -1.7e+308
	var maxFloat64 float64 = +1.7e+308

	fmt.Println(varinfo(minFloat32))
	fmt.Println(varinfo(maxFloat32))
	fmt.Println(varinfo(minFloat64))
	fmt.Println(varinfo(maxFloat64))

	fmt.Printf("%+.38f\n", minFloat32)
	fmt.Printf("%+.38f\n", maxFloat32)

	fmt.Printf("%+.308f\n", minFloat64)
	fmt.Printf("%+.308f\n", maxFloat64)
}
