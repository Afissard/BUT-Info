package pgcd

import "testing"

func TestEgal(t *testing.T) {
	if pgcd(1, 1) != 1 {
		t.Fail()
	}
}

func TestPremier(t *testing.T) {
	if pgcd(3, 5) != 1 {
		t.Fail()
	}
}

func TestDifferent(t *testing.T) {
	if pgcd(220, 315) != 5 {
		t.Fail()
	}
}
