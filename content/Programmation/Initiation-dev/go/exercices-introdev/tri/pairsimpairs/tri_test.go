package tri

import "testing"

func TestVide(t *testing.T) {
	deb := []int{}
	res := []int{}
	tab := []int{}
	tri(tab)
	if !equal(tab, res) {
		t.Error("Une fois trié", deb, "devrait donner", res, "mais donne", tab)
	}
}

func TestExemple(t *testing.T) {
	deb := []int{3, 4, 1, 7, 2, 6, 5}
	res := []int{2, 4, 6, 7, 5, 3, 1}
	tab := []int{3, 4, 1, 7, 2, 6, 5}
	tri(tab)
	if !equal(tab, res) {
		t.Error("Une fois trié", deb, "devrait donner", res, "mais donne", tab)
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
