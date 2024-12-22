package suiteparam

import "testing"

func Test0(t *testing.T) {
	u0 := termeparam(0)
	if u0 != 0 {
		t.Error("termeparam(0) doit valoir 0, mais vous retournez", u0)
	}
}

func TestU5(t *testing.T) {
	u5 := termeparam(5)
	if u5 != 315 {
		t.Error("termeparam(5) doit valoir 315, mais vous retournez", u5)
	}
}

func TestU10(t *testing.T) {
	u10 := termeparam(10)
	if u10 != 20470 {
		t.Error("termeparam(10) doit valoir 20470, mais vous retournez", u10)
	}
}
