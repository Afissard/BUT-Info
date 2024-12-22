package main

func f(x int, y int) {
	x = x + y
}

func main() {
	var n int = 1
	var m int = 2
	f(n, m)
}
