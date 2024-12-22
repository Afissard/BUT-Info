package positifs

import "testing"

func TestVide(t *testing.T) {
	num := positifs([]int{})
	if num != 0 {
		t.Error("Il y a 0 nombres strictement positifs dans le tableau vide, vous en comptez", num)
	}
}

func TestExemple(t *testing.T) {
	num := positifs([]int{-1, -3, 4, 2, 6, -7, 8, -5, 2, 4})
	if num != 6 {
		t.Error("Il y a 6 nombres strictement positifs dans le tabelau [-1 -3 4 2 6 -7 8 -5 2 4], vous en comptez", num)
	}
}

// ajoutés après le test machine

func TestExemple2(t *testing.T) {
	num := positifs([]int{-1, -3, -4, -2, -6, -7, -8, -5, -2, -4})
	if num != 0 {
		t.Fail()
	}
}

func TestExemple3(t *testing.T) {
	num := positifs([]int{0, 1, 2, -1, -2})
	if num != 2 {
		t.Fail()
	}
}
