package puissant

import "testing"

func TestVide(t *testing.T) {
	if !puissant("") {
		t.Error("Le mot vide est puissant (pour toutes les puissances)")
	}
}

func TestCarre(t *testing.T) {
	if !puissant("bonbon") {
		t.Error("Le mot \"bonbon\" est puissant (puissance 2)")
	}
}

func TestCube(t *testing.T) {
	if !puissant("lalala") {
		t.Error("Le mot \"lalala\" est puissant (puissance 3)")
	}
}

func TestFaux(t *testing.T) {
	if puissant("lalalilala") {
		t.Error("Le mot \"lalalilala\" n'est pas puissant")
	}
}

// Ajouté après le premier test machine 2021-2022

func TestVrai1(t *testing.T) {
	// Quelques étudiants pensaient que quand un motif potentiel avait été trouvé
	// c'était forcément ce motif qui se répétait : ici "la" se répète au début,
	// mais le mot n'est pas une puissance de "la", c'est une puissance de "lalali"
	if !puissant("lalalilalali") {
		t.Fail()
	}
}

func TestVrai2(t *testing.T) {
	if !puissant("bb") {
		t.Fail()
	}
}
