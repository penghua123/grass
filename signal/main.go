package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for s := range c {
			switch s {
			case os.Interrupt, os.Kill:
				fmt.Println("退出", s)
				exitFunc()
			default:
				fmt.Println("other", s)
			}
		}
	}()
	fmt.Println("进程启动...")
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
		time.Sleep(time.Second)
	}
}
func exitFunc() {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
}
