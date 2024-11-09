package main

import (
	"fmt"
	"time"
)

func inc(ptrx *int) {
	for {
		(*ptrx)++
		time.Sleep(time.Second)
	}
}

func main() {

	var x int

	go inc(&x)

	for {
		fmt.Println(x)
		time.Sleep(250 * time.Millisecond)
	}

}
