package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ping() {
	for {
		fmt.Println("ping")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func pong() {
	for {
		fmt.Println("pong")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func main() {

	go ping()
	pong()

}
