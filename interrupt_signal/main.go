package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("exit")
			return
		default:
			fmt.Println("running...")
			time.Sleep(3)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	waitForSignal()
	fmt.Println("stopping all job")
	wg.Wait()
}

func waitForSignal() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	fmt.Println(<-sigs)
	os.Exit(1)
}
