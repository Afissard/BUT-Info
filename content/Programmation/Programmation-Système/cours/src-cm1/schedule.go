package main

import (
	"fmt"
	"time"
)

func echo(s string) {
	fmt.Println(s)
}

func main() {

	for {
		go echo("A")
		go echo("B")
		time.Sleep(250 * time.Millisecond)
		fmt.Println("------------")
		time.Sleep(250 * time.Millisecond)
	}

}
