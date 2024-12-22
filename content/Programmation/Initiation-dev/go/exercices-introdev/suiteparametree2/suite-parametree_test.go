package suiteparam

import "testing"

func Test2(t *testing.T) {
	res := termeparam(2)
	if res != 8 {
		t.Error("La suite pour laquelle u(0) vaut 2 est telle que u(2)=8 mais vous retournez", res)
	}
}

func Test11(t *testing.T) {
	res := termeparam(11)
	if res != 487223 {
		t.Error("La suite pour laquelle u(0) vaut 11 est telle que u(11)=8 mais vous retournez", res)
	}
}
