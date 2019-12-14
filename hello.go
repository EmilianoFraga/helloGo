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

    testPrimitiveTypes()
}

// Varinfo displays the value and type info of a variable
func varinfo(x interface{}) string {
    return fmt.Sprintf("var info: %v %T", x, x)
}

// TestPrimitiveTypes just does... test primitive types
func testPrimitiveTypes() {
    var minI8    int8   = -0x80
    var maxI8    int8   =  0x7f
    var minI16   int16  = -0x8000
    var maxI16   int16  =  0x7fff
    var minI32   int32  = -0x80000000
    var maxI32   int32  =  0x7fffffff
    var minI64   int64  = -0x8000000000000000
    var maxI64   int64  =  0x7fffffffffffffff

    var minUI8    uint8   =  0x0
    var maxUI8    uint8   =  0xff
    var minUI16   uint16  =  0x0
    var maxUI16   uint16  =  0xffff
    var minUI32   uint32  =  0x0
    var maxUI32   uint32  =  0xffffffff
    var minUI64   uint64  =  0x0
    var maxUI64   uint64  =  0xffffffffffffffff

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
}
