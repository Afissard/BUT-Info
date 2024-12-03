package main

import "fmt"

func ping(ready, next, done chan bool) {
	for i := 0; i < 5; i++ {
		<-ready
		fmt.Println("Ping")
		next <- true
	}
	done <- true
}

func pong(ready, next, done chan bool) {
	for i := 0; i < 5; i++ {
		<-ready
		fmt.Println("Pong")
		next <- true
	}
	done <- true
}

func main() {
	pingchan := make(chan bool, 1)
	pongchan := make(chan bool, 1)
	done := make(chan bool, 2)

	pingchan <- true // il faut une ball pour jouer (ツ)

	go ping(pingchan, pongchan, done)
	go pong(pongchan, pingchan, done)

	<-done
	<-done
}
