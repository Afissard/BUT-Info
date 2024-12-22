package egalite

import "testing"

func TestVide(t *testing.T) {
	if !egalite([]int{}, []int{}) {
		t.Error("egalite([]int{}, []int{}) devrait retourner true mais retourne false")
	}
	if !egalite(nil, []int{}) {
		t.Error("egalite(nil, []int{}) devrait retourner true mais retourne false")
	}
	if !egalite([]int{}, nil) {
		t.Error("egalite([]int{}, nil) devrait retourner true mais retourne false")
	}
	if !egalite(nil, nil) {
		t.Error("egalite(nil, nil) devrait retourner true mais retourne false")
	}
}

func TestVrai(t *testing.T) {
	if !egalite([]int{1, 2, 3, 4, 5}, []int{2, 4, 3, 1, 5}) {
		t.Error("egalite([]int{1, 2, 3, 4, 5}, []int{2, 4, 3, 1, 5}) devrait retourner true mais retourne false")
	}
}

func TestFaux(t *testing.T) {
	if egalite([]int{1, 2, 3, 4, 5}, []int{1, 4, 3, 6, 5}) {
		t.Error("egalite([]int{1, 2, 3, 4, 5}, []int{1, 4, 3, 6, 5}) devrait retourner false mais retourne true")
	}
}

// Ajouté après le deuxième test machine 2021-2022
func TestLongueur(t *testing.T) {
	if egalite([]int{}, []int{1}) {
		t.Fail()
	}
	if egalite([]int{1}, []int{}) {
		t.Fail()
	}
	if egalite(nil, []int{1}) {
		t.Fail()
	}
	if egalite([]int{1}, nil) {
		t.Fail()
	}
}
