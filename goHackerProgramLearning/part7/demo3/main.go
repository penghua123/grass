package main

import "fmt"
func main() {
	s := []int{1,2,3}
	fmt.Println(s)
	s1 := []int{4,5,6,7,8,9}
	copy(s,s1)
	fmt.Println(s)
	s1[0]=10
	fmt.Println(s)
}
