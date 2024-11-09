package main

import (
	"fmt"
	"time"
)

const maxroutine = 1000000

func echoID(ID int) {
	for {
		fmt.Println(ID)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	for i := 1; i < maxroutine; i++ {
		go echoID(i)
	}

	echoID(maxroutine)
}
