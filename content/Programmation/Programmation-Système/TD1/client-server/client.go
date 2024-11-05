package main

import (
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("unix", "test.sock")
	if err != nil {
		log.Println("Dial error:", err)
		return
	}
	defer conn.Close()
	
	_, err = net.Conn.Write(conn, []byte("hello\n"))
	if err != nil {
		return
	}
	
	log.Println("Je suis connecté")

}
