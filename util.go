package main

import (
	"fmt"
)

// varinfo displays the value and type info of a variable
func varinfo(x interface{}) string {
	return fmt.Sprintf("var info: %v %T", x, x)
}

// varhex displays integer value in hexadecimal format
func varhex(x interface{}) string {
	return fmt.Sprintf("var hex: %32x", x)
}

