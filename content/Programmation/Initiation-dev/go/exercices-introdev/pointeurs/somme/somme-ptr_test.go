package somme

import "testing"

func Test0(t *testing.T) {
	var x int = 5
	somme([]*int{}, &x)
	if x != 0 {
		t.Error("La somme d'un tableau vide doit donner 0 mais donne", x)
	}
}

func TestUnTableau(t *testing.T) {
	var x int
	var a, b, c, d, e, f, g int = 2, 3, 1, 7, 5, 3, 3
	somme([]*int{&a, &b, &c, &d, &e, &f, &g}, &x)
	if x != 24 {
		t.Error("2 + 3 + 1 + 7 + 5 + 3 + 3 vaut 24 mais donne", x)
	}
}
