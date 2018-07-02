package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		c.Write(buf[:n])
	}
}
