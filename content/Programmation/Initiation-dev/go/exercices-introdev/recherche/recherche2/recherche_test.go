package recherche

import "testing"

func TestVide(t *testing.T) {
	trouve, _, _ := recherche([]int{})
	if trouve {
		t.Error("La recherche dans un tableau vide devrait retourner false")
	}
	trouve, _, _ = recherche(nil)
	if trouve {
		t.Error("La recherche dans le tableau nil devrait retourner false")
	}
}

func TestTrouve(t *testing.T) {
	trouve, pos, val := recherche([]int{7, 3, 5, -2, -3, 21})
	if !trouve {
		t.Error("Le tableau [7 3 5 -2 -3 21] contient des entiers strictement positifs et pourtant la recherche retourne false")
	}
	if pos != 1 {
		t.Error("La position du plus petit entier strictement positif dans [7 3 5 -2 -3 21] est 1 mais vous retournez ", pos)
	}
	if val != 3 {
		t.Error("Le plus petit entier strictement positif dans [7 3 5 -2 -3 21] est 3 mais vous retournez ", val)
	}
}

func TestPasTrouve(t *testing.T) {
	var entierParDefaut int
	trouve, pos, val := recherche([]int{-7, -3, -5, -2, -3, -21})
	if trouve {
		t.Error("Le tableau [-7 -3 -5 -2 -3 -21] ne contient pas d'entiers strictement positifs et pourtant la recherche retourne true")
	}
	if pos != entierParDefaut {
		t.Error("Le tableau [-7 -3 -5 -2 -3 -21] ne contient pas d'entiers strictement positifs et pourtant la recherche retourne la position ", pos)
	}
	if val != entierParDefaut {
		t.Error("Le tableau [-7 -3 -5 -2 -3 -21] ne contient pas d'entiers strictement positifs et pourtant la recherche retourne la valeur ", val)
	}
}
