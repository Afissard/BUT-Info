package existe

import "testing"

func TestPresent(t *testing.T) {
	if !existe("../fichiers-tests/fichier1") {
		t.Fail()
	}
}

func TestAbsent(t *testing.T) {
	if existe("cefichiernexistepas") {
		t.Fail()
	}
}
