package calcul

import "testing"

func TestZero(t *testing.T) {
	res := calcul(0)
	if res != 5 {
		t.Error("calcul(0) devrait retourner 5 mais retourne", res)
	}
}

func TestDix(t *testing.T) {
	res := calcul(10)
	if res != 35 {
		t.Error("calcul(10) devrait retourner 35 mais retourne", res)
	}
}

func TestMoinsCinq(t *testing.T) {
	res := calcul(-5)
	if res != -10 {
		t.Error("calcul(-5) devrait retourner -10 mais retourne", res)
	}
}
