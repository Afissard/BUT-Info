package syracuse

import "testing"

func TestU0(t *testing.T) {
	res := syracuse(0, 0)
	if res != 0 {
		t.Error("syracuse(0, 0) retourne ", res, " mais devrait retourner 0")
	}
	res = syracuse(0, 1)
	if res != 1 {
		t.Error("syracuse(0, 1) retourne ", res, " mais devrait retourner 1")
	}
}

func TestU1(t *testing.T) {
	res := syracuse(1, 2)
	if res != 1 {
		t.Error("syracuse(1, 2) retourne ", res, " mais devrait retourner 1")
	}
	res = syracuse(1, 3)
	if res != 10 {
		t.Error("syracuse(1, 3) retourne ", res, " mais devrait retourner 10")
	}
}

func TestExemple(t *testing.T) {
	res := syracuse(10, 27)
	if res != 214 {
		t.Error("syracuse(10, 27) retourne ", res, " mais devrait retourner 214")
	}
}
