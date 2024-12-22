package lignes

import "testing"

func TestPasDeFichier(t *testing.T) {
	if lignes("cefichiernexistepas") != -1 {
		t.Fail()
	}
}

func TestFichierVide(t *testing.T) {
	if lignes("../fichiers-tests/vide") != 0 {
		t.Fail()
	}
}

func TestFichier1(t *testing.T) {
	if lignes("../fichiers-tests/fichier1") != 5 {
		t.Fail()
	}
}

func TestFichier2(t *testing.T) {
	if lignes("../fichiers-tests/fichier2") != 5 {
		t.Fail()
	}
}

func TestFichier3(t *testing.T) {
	if lignes("../fichiers-tests/fichier3") != 50 {
		t.Fail()
	}
}
