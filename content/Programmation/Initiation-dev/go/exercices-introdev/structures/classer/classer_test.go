package classer

import "testing"

func TestVide(t *testing.T) {
	res := classer([]etudiant{})
	attendu := []etudiant{}
	if !egal(res, attendu) {
		t.Error("Le résultat attendu était un tableau vide, ",
			"mais votre fonction retourne ", res)
	}
}

func TestExemple(t *testing.T) {
	res := classer([]etudiant{
		{"Jezequel", "Maël", 12.3},
		{"Berdjugin", "Jean-François", 16.6},
		{"Tamzalit", "Dalila", 15.2},
		{"Hadj-Rabia", "Nassim", 15.2},
		{"Jezequel", "Loïg", 12.3},
	})
	attendu := []etudiant{
		{"Berdjugin", "Jean-François", 16.6},
		{"Hadj-Rabia", "Nassim", 15.2},
		{"Tamzalit", "Dalila", 15.2},
		{"Jezequel", "Loïg", 12.3},
		{"Jezequel", "Maël", 12.3},
	}
	if !egal(res, attendu) {
		t.Error("Le résultat attendu était ", attendu,
			" mais votre fonction retourne ", res)
	}
}

// Tests ajoutés après le troisième test machine 2021-2022

func TestExemple2(t *testing.T) {
	res := classer([]etudiant{
		{"A", "A", 10},
		{"A", "B", 10},
		{"A", "C", 11},
		{"B", "A", 12},
		{"B", "B", 13},
		{"C", "A", 9},
		{"D", "A", 8},
		{"E", "A", 15},
		{"F", "A", 16},
		{"F", "B", 8},
		{"G", "A", 19},
		{"H", "A", 4},
		{"I", "A", 7},
		{"J", "A", 10},
		{"K", "A", 11},
	})
	attendu := []etudiant{
		{"G", "A", 19},
		{"F", "A", 16},
		{"E", "A", 15},
		{"B", "B", 13},
		{"B", "A", 12},
		{"A", "C", 11},
		{"K", "A", 11},
		{"A", "A", 10},
		{"A", "B", 10},
		{"J", "A", 10},
		{"C", "A", 9},
		{"D", "A", 8},
		{"F", "B", 8},
		{"I", "A", 7},
		{"H", "A", 4},
	}
	if !egal(res, attendu) {
		t.Error("Le résultat attendu était ", attendu,
			" mais votre fonction retourne ", res)
	}
}

// fonction utilitaire pour les tests
func egal(t1, t2 []etudiant) bool {
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
