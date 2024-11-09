package main

import (
	"fmt"
	"time"
)

var x int

func inc() {
	for {
		x++
		time.Sleep(time.Second)
	}
}

func main() {

	go inc()

	for {
		fmt.Println(x)
		time.Sleep(250 * time.Millisecond)
	}

}
