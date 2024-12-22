package bougies

import "testing"

func TestBougies1(t *testing.T) {
	res := bougies(23, 5, 1990)
	if res != 4 {
		t.Error("bougies(23, 5, 1990) devrait retourner 4 mais retourne ", res)
	}
}

func TestBougies2(t *testing.T) {
	res := bougies(17, 4, 1990)
	if res != 4 {
		t.Error("bougies(17, 4, 1990) devrait retourner 4 mais retourne ", res)
	}
}

func TestBougies3(t *testing.T) {
	res := bougies(25, 2, 1990)
	if res != 3 {
		t.Error("bougies(25, 2, 1990) devrait retourner 3 mais retourne ", res)
	}
}

// Ajouté après le test 1 de 2022-2023

func TestBougies4(t *testing.T) {
	res := bougies(17, 3, 1986)
	if res != 0 {
		t.Fail()
	}
}

func TestBougies5(t *testing.T) {
	res := bougies(17, 10, 2022)
	if res != 36 {
		t.Fail()
	}
}
