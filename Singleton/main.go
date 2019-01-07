package main

import (
	"fmt"
	"sync"
)

type singleton map[string]string

var instance singleton
var once sync.Once

func New() singleton {
	once.Do(func() {
		instance = make(singleton)
	})
	return instance
}

func main() {
	s := New()
	s["test1"] = "aa"
	fmt.Println(s)

	s1 := New()
	s1["test2"] = "bb"
	fmt.Println(s1)
}
