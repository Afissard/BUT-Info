package main

import (
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("listen error:", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		log.Println("accept error:", err)
		return
	}
	defer conn.Close()
	log.Println("Le client s'est connecté")

	buf := make([]byte, 1024)
    len, err := conn.Read(buf)
    if err != nil {
        log.Printf("Error reading: %#v\n", err)
        return
    }
    log.Printf("Message reçu : %s\n", string(buf[:len]))
	conn.Write([]byte("Message reçu.\n"))

	time.Sleep(10 * time.Second)
}
