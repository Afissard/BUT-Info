package ppcm

import "testing"

func TestEgal1(t *testing.T) {
	res := ppcm([]uint{1, 1})
	if res != 1 {
		t.Fail()
	}
}

func TestEgal2(t *testing.T) {
	res := ppcm([]uint{2, 2, 2, 2, 2})
	if res != 2 {
		t.Fail()
	}
}

func TestNonEgal(t *testing.T) {
	res := ppcm([]uint{12, 9})
	if res != 36 {
		t.Fail()
	}
}

func TestPremiers(t *testing.T) {
	res := ppcm([]uint{2, 3, 5, 7, 11})
	if res != 2310 {
		t.Fail()
	}
}

func TestPlusieurs(t *testing.T) {
	res := ppcm([]uint{5, 12, 3, 8, 9})
	if res != 360 {
		t.Fail()
	}
}
