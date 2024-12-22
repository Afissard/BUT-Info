package monnaie

import (
	"testing"
)

func TestRendu1(t *testing.T) {
	if euros, centimes, _ := rendreMonnaie(5, 50, 5, 70); euros != 0 || centimes != 20 {
		t.Fail()
	}
}

func TestRendu2(t *testing.T) {
	if euros, centimes, _ := rendreMonnaie(5, 50, 6, 70); euros != 1 || centimes != 20 {
		t.Fail()
	}
}

func TestRendu3(t *testing.T) {
	if euros, centimes, _ := rendreMonnaie(5, 50, 6, 30); euros != 0 || centimes != 80 {
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
