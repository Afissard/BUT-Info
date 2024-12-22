package monnaie2

import (
	"testing"
)

func TestRendu1(t *testing.T) {
	euros, centimes, _ := rendreMonnaie(5, 50, 5, 70)
	if somme(euros) != 0 || somme(centimes) != 20 {
		t.Fail()
	}
}

func TestRendu2(t *testing.T) {
	euros, centimes, _ := rendreMonnaie(5, 50, 6, 70)
	if somme(euros) != 1 || somme(centimes) != 20 {
		t.Fail()
	}
}

func TestRendu3(t *testing.T) {
	euros, centimes, _ := rendreMonnaie(5, 50, 6, 30)
	if somme(euros) != 0 || somme(centimes) != 80 {
		t.Fail()
	}
}

func TestRendu4(t *testing.T) {
	euros, centimes, _ := rendreMonnaie(1254, 70, 3127, 50)
	if somme(euros) != 1872 || somme(centimes) != 80 {
		t.Fail()
	}
}

func TestPieces1(t *testing.T) {
	euros, centimes, _ := rendreMonnaie(5, 50, 5, 70)
	if !testPiecesEuros(euros) || !testPiecesCentimes(centimes) {
		t.Fail()
	}
}

func TestPieces2(t *testing.T) {
	euros, centimes, _ := rendreMonnaie(5, 50, 6, 70)
	if !testPiecesEuros(euros) || !testPiecesCentimes(centimes) {
		t.Fail()
	}
}

func TestPieces3(t *testing.T) {
	euros, centimes, _ := rendreMonnaie(5, 50, 6, 30)
	if !testPiecesEuros(euros) || !testPiecesCentimes(centimes) {
		t.Fail()
	}
}

func TestPieces4(t *testing.T) {
	euros, centimes, _ := rendreMonnaie(3127, 50, 1254, 70)
	if !testPiecesEuros(euros) || !testPiecesCentimes(centimes) {
		t.Fail()
	}
}

func TestError1(t *testing.T) {
	if _, _, err := rendreMonnaie(5, 50, 2, 20); err != errPasAssezPaye {
		t.Fail()
	}
}

func TestError2(t *testing.T) {
	if _, _, err := rendreMonnaie(5, 50, 3, 70); err != errPasAssezPaye {
		t.Fail()
	}
}

func TestError3(t *testing.T) {
	if _, _, err := rendreMonnaie(5, 50, 5, 20); err != errPasAssezPaye {
		t.Fail()
	}
}

func TestError4(t *testing.T) {
	if _, _, err := rendreMonnaie(5, 50, 6, 70); err != nil {
		t.Fail()
	}
}

func TestError5(t *testing.T) {
	if _, _, err := rendreMonnaie(5, 50, 6, 20); err != nil {
		t.Fail()
	}
}

func TestError6(t *testing.T) {
	if _, _, err := rendreMonnaie(5, 50, 5, 50); err != nil {
		t.Fail()
	}
}

// Fonctions utilitaires pour les tests
func somme(pieces []int) int {
	var s int
	for _, v := range pieces {
		s += v
	}
	return s
}

func testPiecesCentimes(pieces []int) bool {
	for _, v := range pieces {
		if v != 1 && v != 2 && v != 5 && v != 10 && v != 20 && v != 50 {
			return false
		}
	}
	return true
}

func testPiecesEuros(pieces []int) bool {
	for _, v := range pieces {
		if v != 1 && v != 2 && v != 5 && v != 10 && v != 20 && v != 50 &&
			v != 100 && v != 200 && v != 500 {
			return false
		}
	}
	return true
}
