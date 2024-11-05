package main

import (
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "172.21.180.254:8080")
	if err != nil {
		log.Println("Dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("Je suis connecté")

	_, err = net.Conn.Write(conn, []byte("hello there\n"))
	if err != nil {
		return
	}

	// réception d'un msg
	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	if err != nil {
		log.Printf("Error reading: %#v\n", err)
		return
	}
	log.Printf("Message reçu : %s\n", string(buf[:len]))
}
