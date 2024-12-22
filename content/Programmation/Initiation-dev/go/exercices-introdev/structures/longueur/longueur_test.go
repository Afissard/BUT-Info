package longueur

import "testing"

func TestVide(t *testing.T) {
	res := longueur(chaine{})
	if res != 0 {
		t.Error("La longueur d'une chaîne vide doit être 0, vous retournez ", res)
	}
}

func TestExemple(t *testing.T) {
	e5 := element{}
	e4 := element{next: &e5}
	e3 := element{next: &e4}
	e2 := element{next: &e3}
	e1 := element{next: &e2}
	res := longueur(chaine{debut: &e1})
	if res != 5 {
		t.Error("La longueur d'une chaîne de 5 éléments doit être 5, vous retournez ", res)
	}
}
