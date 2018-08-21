package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	queue := make(chan int, 100)
	signal := make(chan int)
	go reader(queue, signal)
	writer(queue, signal)
	time.Sleep(100000)
}

func reader(queue chan<- int, signal chan<- int) {
	docNum := 0
	for i := 0; i < 1000; i++ {
		queue <- i
		fmt.Println(i)
		docNum++
		signal <- docNum
	}
	close(queue)
	close(signal)
	fmt.Println("Send document is", docNum)
}

func writer(queue <-chan int, signal <-chan int) {
	docNum := 0
	j := 0
	var wg sync.WaitGroup
	threads := 100
	for j < threads {

		_, ok := <-signal
		if !ok {
			fmt.Println("Received the stop signal")
			wg.Wait()
			break
		}
		wg.Add(1)
		go func(queue <-chan int) {
			defer wg.Done()
			v := <-queue
			fmt.Println("receive:", v)
		}(queue)
		docNum++
		j++
		if j >= threads {
			wg.Wait()
			j = 0
		}
		//if docNum > 999 {
		//	break
		//}
		//_, ok := <-signal
		//if sig == docNum
	}
	fmt.Println("Received document is", docNum)
}
