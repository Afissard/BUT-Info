package suitehist

import "testing"

func TestU0(t *testing.T) {
	u0 := terme(0)
	if u0 != 0 {
		t.Error("u0 doit valoir 0, mais vous retournez", u0)
	}
}

func TestU5(t *testing.T) {
	u5 := terme(5)
	if u5 != 6 {
		t.Error("u5 doit valoir 6, mais vous retournez", u5)
	}
}

func TestU10(t *testing.T) {
	u10 := terme(10)
	if u10 != 14 {
		t.Error("u10 doit valoir 14, mais vous retournez", u10)
	}
}

// Ajouté après le test

func TestU1(t *testing.T) {
	u1 := terme(1)
	if u1 != 2 {
		t.Error("u1 doit valoir 2, mais vous retournez", u1)
	}
}

func TestU2(t *testing.T) {
	u2 := terme(2)
	if u2 != 2 {
		t.Error("u2 doit valoir 2, mais vous retournez", u2)
	}
}
