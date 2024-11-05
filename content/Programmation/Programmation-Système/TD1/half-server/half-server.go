package main

import (
	"log"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("unix", "test.sock")
	if err != nil {
		log.Println("listen error:", err)
		return
	}
	defer listener.Close()

	time.Sleep(10 * time.Second)

}
