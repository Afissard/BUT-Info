package nombre

import "testing"

func TestPresent(t *testing.T) {
	if !nombre("../fichiers-tests/contient1.txt") {
		t.Error("Le fichier contient1.txt contient le chiffre 1 mais votre fonction indique que ce n'est pas le cas")
	}
}

func TestAbsent(t *testing.T) {
	if nombre("../fichiers-tests/contientpas.txt") {
		t.Error("Le fichier contientpas1.txt ne contient pas le chiffre 1 mais votre fonction indique que si")
	}
}
