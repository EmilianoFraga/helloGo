package main

import "fmt"

func testArrays() {
	arTen := [...] int{ 0, 1, 2, 3, 4 ,5 ,6 ,7 ,8 ,9 }
	fmt.Println(varinfo(arTen))

	var arFive [5] int
	arFive = [5] int{ 0, 1, 2, 3, 4 }
	fmt.Println(varinfo(arFive))

	var arFour = [...] int{ 3: 20, 0: 13 }
	fmt.Println(varinfo(arFour))

	var arThree [3] string
	arThree[0] = "one"
	arThree[1] = "two"
	arThree[2] = "three"
	fmt.Println(varinfo(arThree))
}
