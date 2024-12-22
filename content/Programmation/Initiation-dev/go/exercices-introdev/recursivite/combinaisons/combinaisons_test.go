package combinaison

import (
	"sort"
	"testing"
)

func TestVide(t *testing.T) {
	num := len(combinaisons(0))
	if num != 0 {
		t.Error("Il n'y a aucune combinaison d'entiers entre 1 et 0 dont la somme est 0, mais vous en retournez", num)
	}
}

func Test3(t *testing.T) {
	expected := [][]int{
		{1, 1, 1},
		{1, 2},
		{3},
	}
	obtained := combinaisons(3)
	if !equal(expected, obtained) {
		t.Error("combinaisons(3) devrait retourner", expected, "mais retourne", obtained)
	}
}

func Test5(t *testing.T) {
	expected := [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 2},
		{1, 2, 2},
		{1, 1, 3},
		{2, 3},
		{1, 4},
		{5},
	}
	obtained := combinaisons(5)
	if !equal(expected, obtained) {
		t.Error("combinaisons(5) devrait retourner", expected, "mais retourne", obtained)
	}
}

// Fonction auxiliaire pour tester que les ensembles de combinaisons sont les mêmes

func equal(e1, e2 [][]int) bool {
	if len(e1) != len(e2) {
		return false
	}
	for i := 0; i < len(e1); i++ {
		// on trie toutes les combinaisons
		sort.Slice(e1[i], func(ii, jj int) bool { return e1[i][ii] < e1[i][jj] })
		sort.Slice(e2[i], func(ii, jj int) bool { return e2[i][ii] < e2[i][jj] })
	}
	sort.Slice(e1, lesscomb(e1))
	sort.Slice(e2, lesscomb(e2))
	for i := 0; i < len(e1); i++ {
		if !equalcomb(e1[i], e2[i]) {
			return false
		}
	}
	return true
}

func lesscomb(e [][]int) func(int, int) bool {
	// les combinaisons sont supposées triées
	return func(i, j int) bool {
		if len(e[i]) != len(e[j]) {
			return len(e[i]) < len(e[j])
		}
		for ii, v := range e[i] {
			if v != e[j][ii] {
				return v < e[j][ii]
			}
		}
		return true
	}
}

func equalcomb(c1, c2 []int) bool {
	// les combinaisons sont supposées triées
	if len(c1) != len(c2) {
		return false
	}
	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			return false
		}
	}
	return true
}
