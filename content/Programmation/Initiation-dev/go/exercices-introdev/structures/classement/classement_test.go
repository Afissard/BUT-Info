package classement

import "testing"

func TestExemple(t *testing.T) {
	res := classement([]etudiant{
		{"Berdjugin", "Jean-François", 16.6},
		{"Hadj-Rabia", "Nassim", 15.2},
		{"Tamzalit", "Dalila", 15.2},
		{"Jezequel", "Loïg", 12.3},
		{"Jezequel", "Maël", 12.3},
	})
	attendu := "1. Berdjugin Jean-François\n" +
		"2. Hadj-Rabia Nassim\n" +
		"2. Tamzalit Dalila\n" +
		"4. Jezequel Loïg\n" +
		"4. Jezequel Maël\n"
	if res != attendu {
		t.Error("Votre affichage donne :\n", res,
			"\nEt on attendait :\n", attendu)
	}
}

func TestVide(t *testing.T) {
	res := classement([]etudiant{})
	attendu := ""
	if res != attendu {
		t.Error("Sur un tableau vide, votre programme retourne '", res, "'")
	}
}

// Ajouté après le troisième test machine 2021-2022

func TestExemple2(t *testing.T) {
	res := classement([]etudiant{
		{"G", "A", 19},
		{"F", "A", 16},
		{"E", "A", 15},
		{"B", "B", 13},
		{"B", "A", 12},
		{"A", "C", 11},
		{"K", "A", 11},
		{"A", "A", 10},
		{"A", "B", 10},
		{"J", "A", 10},
		{"C", "A", 9},
		{"D", "A", 8},
		{"F", "B", 8},
		{"I", "A", 7},
		{"H", "A", 4},
	})
	attendu := "1. G A\n" +
		"2. F A\n" +
		"3. E A\n" +
		"4. B B\n" +
		"5. B A\n" +
		"6. A C\n" +
		"6. K A\n" +
		"8. A A\n" +
		"8. A B\n" +
		"8. J A\n" +
		"11. C A\n" +
		"12. D A\n" +
		"12. F B\n" +
		"14. I A\n" +
		"15. H A\n"
	if res != attendu {
		t.Error("Votre affichage donne :\n", res,
			"\nEt on attendait :\n", attendu)
	}
}
