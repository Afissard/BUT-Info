package main

import "fmt"

var turn1 bool

func routine1() {
	for {
		if turn1 {
			fmt.Println("Routine 1")
			turn1 = false
		}
	}
}

func routine2() {
	for {
		if !turn1 {
			fmt.Println("Routine 2")
			turn1 = true
		}
	}
}

func routine1bis(t, f chan bool) {
	for {
		if <-t {
			fmt.Println("Routine 1")
			f <- false
		}
	}
}

func routine2bis(t, f chan bool) {
	for {
		if !<-f {
			fmt.Println("Routine 2")
			t <- true
		}
	}
}


func main() {
	go routine1()
	routine2()

	// var t chan bool = make(chan bool, 1)
	// var f chan bool = make(chan bool, 1)
	// t <- true
	// f <- false
	// go routine1bis(t, f)
	// routine2bis(t, t)
}
