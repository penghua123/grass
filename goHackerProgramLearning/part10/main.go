package main

import "fmt"
func main(){
	ch := make(chan []int, 1)
	s1 := []int{1, 2, 3}
	ch <- s1
	s2 := <-ch

	s2[0] = 100
	fmt.Println(s1, s2)


	ch2 := make(chan [3]int, 1)
	s3 := [3]int{1, 2, 3}
	ch2 <- s3
	s4 := <-ch2

	s3[0] = 100
	fmt.Println(s3, s4)
}
