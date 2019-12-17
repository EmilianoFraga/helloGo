package main

import "fmt"

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
