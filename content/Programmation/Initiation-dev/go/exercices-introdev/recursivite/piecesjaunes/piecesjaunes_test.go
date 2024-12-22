package piecesjaunes

import "testing"

func TestVide(t *testing.T) {
	_, err := meilleurDecomposition(1, []int{})
	if err != errImpossible {
		t.Fail()
	}
}

func TestNulVide(t *testing.T) {
	v, err := meilleurDecomposition(0, []int{})
	if err != nil || v == nil {
		t.Fail()
	}
}

func TestNulVal(t *testing.T) {
	v, _ := meilleurDecomposition(0, []int{1})
	if len(v) != 0 || v == nil {
		t.Fail()
	}
}

func TestNulErr(t *testing.T) {
	_, err := meilleurDecomposition(0, []int{1})
	if err != nil {
		t.Fail()
	}
}

func TestVal(t *testing.T) {
	v, _ := meilleurDecomposition(15, []int{2, 3, 5, 5, 5})
	if len(v) != 3 ||
		mauvaisesValeurs(v, []int{2, 3, 5, 5, 5}) ||
		mauvaiseSomme(v, 15) {
		t.Fail()
	}
}

func TestErr1(t *testing.T) {
	_, err := meilleurDecomposition(15, []int{2, 3, 5, 5, 5})
	if err != nil {
		t.Fail()
	}
}

func TestErr2(t *testing.T) {
	_, err := meilleurDecomposition(15, []int{2, 3, 5})
	if err != errImpossible {
		t.Fail()
	}
}

// Fonctions utilitaires pour les tests
func mauvaisesValeurs(t []int, pieces []int) bool {
principale:
	for _, v := range t {
		for _, vv := range pieces {
			if v == vv {
				continue principale
			}
		}
		return true
	}
	return false
}

func mauvaiseSomme(t []int, s int) bool {
	for _, v := range t {
		s -= v
	}
	return s != 0
}
