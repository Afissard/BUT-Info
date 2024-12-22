package suite

import "testing"

func TestZero(t *testing.T) {
	res := suite(0)
	if res != 0 {
		t.Error("suite(0) devrait retourner 0 mais retourne", res)
	}
}

func TestUn(t *testing.T) {
	res := suite(1)
	if res != 5 {
		t.Error("suite(1) devrait retourner 5 mais retourne", res)
	}
}

func TestCinq(t *testing.T) {
	res := suite(5)
	if res != 605 {
		t.Error("suite(5) devrait retourner 605 mais retourne", res)
	}
}
