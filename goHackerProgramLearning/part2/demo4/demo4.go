package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	//flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	//flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	//flag.Parse()
	//others := cmdLine.Args()
	cmdLine.Parse(os.Args[1:])
	others := cmdLine.Args()
	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("Others are, %s!\n", others)
}
