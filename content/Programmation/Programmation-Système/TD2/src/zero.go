package main

import (
	"fmt"
	"time"
)

func writer(c chan bool) {
	for {
		time.Sleep(500 * time.Millisecond)
		c <- true
		fmt.Println("Écriture")
	}
}

func reader(c chan bool) {
	for {
		time.Sleep(time.Second)
		<-c
		fmt.Println("Lecture")
	}
}

func main() {
	// var c chan bool = make(chan bool)
	var c chan bool = make(chan bool, 32)
	go writer(c)
	reader(c)
}
