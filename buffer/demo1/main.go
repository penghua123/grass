package main

import (
    "bytes"
    "fmt"
)

func main() {
    var buffer bytes.Buffer
    for i := 0; i < 1000; i++ {
        buffer.WriteString("a")
    }
	fmt.Println(buffer.String())
	buffer.ReadString
}
