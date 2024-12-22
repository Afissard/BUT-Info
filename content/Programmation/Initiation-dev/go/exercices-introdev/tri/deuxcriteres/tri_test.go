package tri

import "testing"

func TestVide(t *testing.T) {
	res := []int{}
	tab := []int{}
	tri(tab)
	if !equal(tab, res) {
		t.Error("Trier un tableau vide doit donner un tableau vide, vous retournez", tab)
	}
}

func TestExemple(t *testing.T) {
	res := []int{-1, -2, -6, 3, 4, 5, 7}
	tab := []int{3, 4, -1, 7, -2, -6, 5}
	tri(tab)
	if !equal(tab, res) {
		t.Error("Trier le tableau [3 4 -1 7 -2 -6 5] doit donner [-1, -2, -6, 3, 4, 5, 7] mais vous retournez", tab)
	}
}

// Ajoutés après le test machine

func TestExemple2(t *testing.T) {
	res := []int{1}
	tab := []int{1}
	tri(tab)
	if !equal(tab, res) {
		t.Fail()
	}
}

func TestExemple3(t *testing.T) {
	res := []int{-3, -4, -5, -7, 1, 2, 6}
	tab := []int{-3, -4, 1, -7, 2, 6, -5}
	tri(tab)
	if !equal(tab, res) {
		t.Fail()
	}
}

// Fonction pour vérifier l'égalité de deux tableaux

func equal(t1, t2 []int) bool {
	if len(t1) != len(t2) {
		return false
	}
	for i := 0; i < len(t1); i++ {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}
