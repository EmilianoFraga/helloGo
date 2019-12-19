package main

import (
	"fmt"
	"strconv"
)

func testArrays() {
	var arTen = [...] int{ 0, 1, 2, 3, 4 ,5 ,6 ,7 ,8 ,9 }
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

	var arThreeCopy = arThree
	var arThreeRef = &arThree

	fmt.Println(varinfo(arThree))
	fmt.Println(varinfo(arThreeCopy))
	fmt.Println(varinfo(arThreeRef))

	arThreeCopy[1] = "abracadabra"
	arThreeRef[1] = "mingau"

	fmt.Println(varinfo(arThree))
	fmt.Println(varinfo(arThreeCopy))
	fmt.Println(varinfo(arThreeRef))
}


func testSlices() {
	var arBack = [...] int{ 0, 1, 2, 3, 4 ,5 ,6 ,7 ,8 ,9 }
	var sliceFront []int = arBack[:]
	anotherSlice := arBack[:]

	fmt.Println(varinfo(arBack))
	fmt.Println(varinfo(sliceFront))
	fmt.Println(varinfo(anotherSlice))

	sliceFront[2] = 2123
	fmt.Println(varinfo(arBack))
	fmt.Println(varinfo(sliceFront))
	fmt.Println(varinfo(anotherSlice))

	anotherSlice = sliceFront[3:]
	fmt.Println(varinfo(arBack))
	fmt.Println(varinfo(sliceFront))
	fmt.Println(varinfo(anotherSlice))

	for i, v := range arBack {
		fmt.Printf("arBack[%d] = %d\n", i, v)
	}

	s := ""
	for _, v := range arBack {
		if len(s) > 0 {
			s += ", "
		}
		s += strconv.Itoa(v)
	}
	s = "[" + s
	s += "]"
	fmt.Printf("arBack = %s\n", s)

	s = ""
	for i := range arBack {
		if len(s) > 0 {
			s += ", "
		}
		s += strconv.Itoa(arBack[i])
	}
	s = "[" + s
	s += "]"
	fmt.Printf("arBack = %s\n", s)

    veggies := []string{ "potatoes", "tomatoes", "brinjal" }
    fruits := []string{ "oranges", "apples" }
    food := append( veggies, fruits... )
	fmt.Println("food:", food)
}
