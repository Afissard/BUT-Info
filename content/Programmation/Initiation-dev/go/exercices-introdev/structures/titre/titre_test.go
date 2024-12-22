package titre

import "testing"

func TestLivre(t *testing.T) {
	l := livre{titre: "L'empire ultime", prix: 10.90, numPages: 928}
	titre := obtenirTitre(l)
	if titre != "L'empire ultime" {
		t.Error("Vous trouvez le titre", titre, "alors que ça devrait ếtre l'Empire ultime")
	}
}

// ajouté après le test machine

func TestLivre2(t *testing.T) {
	l := livre{titre: "Un autre titre", prix: 10.90, numPages: 928}
	titre := obtenirTitre(l)
	if titre != "Un autre titre" {
		t.Fail()
	}
}
