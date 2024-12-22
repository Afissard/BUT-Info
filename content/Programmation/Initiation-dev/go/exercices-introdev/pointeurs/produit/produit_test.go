package produit

import "testing"

func Test0(t *testing.T) {
	var x int = 5
	produit([]*int{}, &x)
	if x != 1 {
		t.Error("Le produit des éléments d'un tableau vide devrait être 1, mais vous retournez", x)
	}
}

func TestUnTableau(t *testing.T) {
	var x int
	var a, b, c, d, e, f, g int = 2, 3, 1, 7, 5, 3, 3
	produit([]*int{&a, &b, &c, &d, &e, &f, &g}, &x)
	if x != 1890 {
		t.Error("Le produit de 2, 3, 1, 7, 5, 3 et 3 devrait être 1890 mais vous retournez", x)
	}
}

// ajouté après le test machine

func TestUnTableau2(t *testing.T) {
	var x int
	var a, b, c, d, e, f, g int = -2, 3, 1, 7, 5, 3, 3
	produit([]*int{&a, &b, &c, &d, &e, &f, &g}, &x)
	if x != -1890 {
		t.Fail()
	}
	if a != -2 || b != 3 || c != 1 || d != 7 || e != 5 || f != 3 || g != 3 {
		t.Fail()
	}
}
