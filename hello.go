package main

import "fmt"

func main() {
    i := 345
    fmt.Println("Hello Go")
    s := fmt.Sprintf("This is a formatted output %05d.", i)
    fmt.Println(s)
    fmt.Println(varinfo(i))
    fmt.Println(varinfo(s))
}

func varinfo(x interface{}) string {
    return fmt.Sprintf("varinfo: %v %T", x, x)
}