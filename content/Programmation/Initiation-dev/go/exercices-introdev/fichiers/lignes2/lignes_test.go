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
	if lignes("../fichiers-tests/fichier2") != 3 {
		t.Fail()
	}
}

func TestFichier3(t *testing.T) {
	if lignes("../fichiers-tests/fichier3") != 37 {
		t.Fail()
	}
}

func TestFichier4(t *testing.T) {
	if lignes("../fichiers-tests/fichier4") != 3 {
		t.Fail()
	}
}
