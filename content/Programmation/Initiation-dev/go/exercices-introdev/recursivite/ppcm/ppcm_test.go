package ppcm

import "testing"

func TestEgal(t *testing.T) {
	res := ppcm(1, 1)
	if res != 1 {
		t.Error("ppcm(1, 1) retourne ", res, " mais devrait retourner 1")
	}
}

func TestNonEgal(t *testing.T) {
	res := ppcm(12, 9)
	if res != 36 {
		t.Error("ppcm(12, 9) retourne ", res, " mais devrait retourner 36")
	}
}
