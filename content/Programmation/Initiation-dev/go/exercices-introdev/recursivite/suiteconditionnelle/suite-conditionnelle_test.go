package suitecond

import "testing"

func TestU0(t *testing.T) {
	u0 := terme(0)
	if u0 != 2 {
		t.Error("u0 doit valoir 2, mais vous retournez", u0)
	}
}

func TestU5(t *testing.T) {
	u5 := terme(5)
	if u5 != 95 {
		t.Error("u5 doit valoir 95, mais vous retournez", u5)
	}
}

func TestU10(t *testing.T) {
	u10 := terme(10)
	if u10 != 167 {
		t.Error("u10 doit valoir 167, mais vous retournez", u10)
	}
}
