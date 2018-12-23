package main

import (
	"flag"
	"fmt"
)

var name string

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}
