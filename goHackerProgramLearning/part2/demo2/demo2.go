package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}
