package main

func RetVal() int {
	return 42
}

func AddSub(a, b int) (add, sub int) {
	return a + b, a - b
}

func CallAddSub() (a, b int) {
	a, b = AddSub(4, 5)
	return a, b
}

func RetPtr() *int {
	x := 42
	return &x
}

func ManyArgs(a, b, c, d, e, f, g, h, i, j, k, l int) (r int) {
	r = a*b + c*d - e*f*g + h + i*j*k - l
	return
}

func main() {
	RetVal()
	RetPtr()
	AddSub(3, 3)
	CallAddSub()
	ManyArgs(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
}
