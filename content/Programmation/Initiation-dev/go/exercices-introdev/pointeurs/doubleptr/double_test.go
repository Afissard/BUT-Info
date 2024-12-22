package double

import "testing"

func TestNil1(t *testing.T) {
	var n int
	if double(&n, nil) == nil {
		t.Error("double(&n, nil) devrait retourner une erreur")
	}
}

func TestNil2(t *testing.T) {
	var ptr *int
	if double(nil, &ptr) != nil {
		t.Error("double(nil, &ptr) ne devrait pas retourner d'erreur")
	}
	if ptr != nil {
		t.Error("double(nil, &ptr) devrait mettre nil dans ptr")
	}
}

func TestSuccess(t *testing.T) {
	var n int
	var ptr *int
	if double(&n, &ptr) != nil {
		t.Error("double(&n, &ptr) ne devrait pas retourner d'erreur")
	}
	if ptr != &n {
		t.Error("double(&n, &ptr) devrait mettre l'adresse de n dans ptr")
	}
}
