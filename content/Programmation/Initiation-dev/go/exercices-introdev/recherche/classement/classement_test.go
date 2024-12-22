package classement

import "testing"

func Test(t *testing.T) {
	pairs, petits, grands := classer([]int{})
	if len(pairs) != 0 || len(petits) != 0 || len(grands) != 0 {
		t.Fail()
	}
}

func TestPairs1(t *testing.T) {
	pairs, _, _ := classer([]int{2, 127, 103, 100, 212, 4})
	if len(pairs) != 4 {
		t.Fail()
	} else if !estDans(2, pairs) ||
		!estDans(100, pairs) ||
		!estDans(212, pairs) ||
		!estDans(4, pairs) {
		t.Fail()
	}
}

func TestPairs2(t *testing.T) {
	pairs, _, _ := classer([]int{1, 127, 103, 101, 211, 7})
	if len(pairs) != 0 {
		t.Fail()
	}
}

func TestPetits1(t *testing.T) {
	_, petits, _ := classer([]int{2, 3, 103, 100, 212, 4, 5, 91})
	if len(petits) != 3 {
		t.Fail()
	} else if !estDans(3, petits) ||
		!estDans(5, petits) ||
		!estDans(91, petits) {
		t.Fail()
	}
}

func TestPetits2(t *testing.T) {
	_, petits, _ := classer([]int{100, 12, 10, 101, 211, 70})
	if len(petits) != 0 {
		t.Fail()
	}
}

func TestGrands1(t *testing.T) {
	_, _, grands := classer([]int{2, 3, 103, 100, 212, 4, 5, 91})
	if len(grands) != 1 {
		t.Fail()
	} else if !estDans(103, grands) {
		t.Fail()
	}
}

func TestGrands2(t *testing.T) {
	_, _, grands := classer([]int{100, 12, 10, 108, 218, 70})
	if len(grands) != 0 {
		t.Fail()
	}
}

// Fonctions utilitaires pour les tests
func estDans(v int, t []int) bool {
	for _, vv := range t {
		if v == vv {
			return true
		}
	}
	return false
}
