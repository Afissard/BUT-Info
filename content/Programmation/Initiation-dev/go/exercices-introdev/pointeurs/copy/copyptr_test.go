package copyptr

import "testing"

func TestUn(t *testing.T) {
	var y int
	copyptr(1, []*int{&y})
	if y != 1 {
		t.Error("Echec de la première copie")
	}
	copyptr(2, []*int{&y})
	if y != 2 {
		t.Error("Echec de la deuxième copie")
	}
}

func TestPlusieurs(t *testing.T) {
	var x, y, z int
	copyptr(3, []*int{&x, &y, &z})
	if x != 3 {
		t.Error("Echec de la copie dans x")
	}
	if y != 3 {
		t.Error("Echec de la copie dans y")
	}
	if z != 3 {
		t.Error("Echec de la copie dans z")
	}
}
