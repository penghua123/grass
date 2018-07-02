package main

import (
	"log"
	"net"
)

func main() {
	log.Println("begin dial...")
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Println("dial error:", err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("hello"))
	log.Println("dial ok")
}
