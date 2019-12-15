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
}

// Varinfo displays the value and type info of a variable
func varinfo(x interface{}) string {
	return fmt.Sprintf("var info: %v %T", x, x)
}

// Varhex displays integer value in hexadecimal format
func varhex(x interface{}) string {
	return fmt.Sprintf("var hex: %32x", x)
}

// TestPrimitiveTypes just does... test primitive types
func testIntegerPrimitiveTypes() {
	var minI8 int8 = -0x80
	var maxI8 int8 = 0x7f
	var minI16 int16 = -0x8000
	var maxI16 int16 = 0x7fff
	var minI32 int32 = -0x80000000
	var maxI32 int32 = 0x7fffffff
	var minI64 int64 = -0x8000000000000000
	var maxI64 int64 = 0x7fffffffffffffff

	var minUI8 uint8 = 0x0
	var maxUI8 uint8 = 0xff
	var minUI16 uint16 = 0x0
	var maxUI16 uint16 = 0xffff
	var minUI32 uint32 = 0x0
	var maxUI32 uint32 = 0xffffffff
	var minUI64 uint64 = 0x0
	var maxUI64 uint64 = 0xffffffffffffffff

	fmt.Println(varinfo(minI8))
	fmt.Println(varinfo(maxI8))
	fmt.Println(varinfo(minI16))
	fmt.Println(varinfo(maxI16))
	fmt.Println(varinfo(minI32))
	fmt.Println(varinfo(maxI32))
	fmt.Println(varinfo(minI64))
	fmt.Println(varinfo(maxI64))

	fmt.Println(varinfo(minUI8))
	fmt.Println(varinfo(maxUI8))
	fmt.Println(varinfo(minUI16))
	fmt.Println(varinfo(maxUI16))
	fmt.Println(varinfo(minUI32))
	fmt.Println(varinfo(maxUI32))
	fmt.Println(varinfo(minUI64))
	fmt.Println(varinfo(maxUI64))

	fmt.Println(varhex(minI8))
	fmt.Println(varhex(maxI8))
	fmt.Println(varhex(minI16))
	fmt.Println(varhex(maxI16))
	fmt.Println(varhex(minI32))
	fmt.Println(varhex(maxI32))
	fmt.Println(varhex(minI64))
	fmt.Println(varhex(maxI64))

	fmt.Println(varhex(minUI8))
	fmt.Println(varhex(maxUI8))
	fmt.Println(varhex(minUI16))
	fmt.Println(varhex(maxUI16))
	fmt.Println(varhex(minUI32))
	fmt.Println(varhex(maxUI32))
	fmt.Println(varhex(minUI64))
	fmt.Println(varhex(maxUI64))

	minI8 = -1
	// output is:
	// var hex:                               -1
	// shouldn't be ?
	// var hex:                             0xff
	fmt.Println(varhex(minI8))

	var tempI8 int8 = maxI8 + maxI8
	fmt.Println(varinfo(tempI8))

	var tempUI8 uint8 = maxUI8 + maxUI8
	fmt.Println(varinfo(tempUI8))

	var tempI16 int16 = maxI16 + maxI16
	fmt.Println(varinfo(tempI16))

	var tempUI16 uint16 = maxUI16 + maxUI16
	fmt.Println(varinfo(tempUI16))

	var tempI32 int32 = maxI32 + maxI32
	fmt.Println(varinfo(tempI32))

	var tempUI32 uint32 = maxUI32 + maxUI32
	fmt.Println(varinfo(tempUI32))

	var tempI64 int64 = maxI64 + maxI64
	fmt.Println(varinfo(tempI64))

	var tempUI64 uint64 = maxUI64 + maxUI64
	fmt.Println(varinfo(tempUI64))

	var byteIsUInt8 = maxUI8
	fmt.Println(varinfo(byteIsUInt8))
}

// TestFloatPrimitiveTypes just does... test float primitive types
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

func testConst() {
	const impliedInt32Const1 = 0x7f
	const impliedInt32Const2 = 0x7fff
	const impliedInt32Const3 = 0x7fffffff
	const impliedInt64Const1 = 0x7fffffffffffffff

	// constant 18446744073709551615 overflows int
	// const impliedInt64Const2 = 0xffffffffffffffff
	const impliedInt64Const2 = 0x7fffffffffffffff

	// constant 18446744073709551615 overflows int64
	// const explicitInt64Const int64 = 0xffffffffffffffff
	const explicitInt64Const int64 = -1

	fmt.Println(varinfo(impliedInt32Const1))
	fmt.Println(varinfo(impliedInt32Const2))
	fmt.Println(varinfo(impliedInt32Const3))
	fmt.Println(varinfo(impliedInt64Const1))
	fmt.Println(varinfo(impliedInt64Const2))
	fmt.Println(varinfo(explicitInt64Const))
	fmt.Printf("%x\n", explicitInt64Const)

	const (
		sunday = iota + 1
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
	)

	fmt.Println(varinfo(sunday))
	fmt.Println(varinfo(monday))
	fmt.Println(varinfo(tuesday))
	fmt.Println(varinfo(wednesday))
	fmt.Println(varinfo(thursday))
	fmt.Println(varinfo(friday))
	fmt.Println(varinfo(saturday))
}
