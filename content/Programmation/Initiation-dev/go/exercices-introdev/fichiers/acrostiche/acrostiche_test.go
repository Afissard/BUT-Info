package acrostiche

import "testing"

func TestFichier1(t *testing.T) {
	if acrostiche("../fichiers-tests/fichier1") != "lllll" {
		t.Fail()
	}
}

func TestFichier2(t *testing.T) {
	if acrostiche("../fichiers-tests/fichier2") != "lll" {
		t.Fail()
	}
}

func TestBonjour(t *testing.T) {
	if acrostiche("../fichiers-tests/bonjour") != "Bonjour" {
		t.Fail()
	}
}
