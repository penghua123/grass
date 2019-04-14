package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var uselessChan = make(chan<- int, 1)

	var anotherUselessChan = make(<-chan int, 1)

	fmt.Printf("The useless channels: %v, %v\n",
		uselessChan, anotherUselessChan)

	intChan1 := make(chan int, 3)
	SendInt(intChan1)

	intChan2 := getIntChan()
	for elem := range intChan2 {
		fmt.Printf("The element in intChan2: %v\n", elem)
	}

	_ = GetIntChan(getIntChan)
}


func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}

type Notifier interface {
	SendInt(ch chan<- int)
}

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

type GetIntChan func() <-chan int
