package main

func main() {
	a, b := 1, 2
	_ = additionBasique(a, b)
	_ = additionVersPointeur(a, b)
	_ = additionDepuisPointeurs(&a, &b)
}

//go:noinline
func additionBasique(x, y int) int {
	z := x + y
	return z
}

//go:noinline
func additionVersPointeur(x, y int) *int {
	z := x + y
	return &z
}

//go:noinline
func additionDepuisPointeurs(x, y *int) int {
	z := *x + *y
	return z
}
