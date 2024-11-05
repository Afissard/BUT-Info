package main

import (
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("Dial error:", err)
		return
	}
	defer conn.Close()
	log.Println("Je suis connecté")

	// envoi d'un msg
	_, err = net.Conn.Write(conn, []byte("hello\n"))
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
