package bienrange

import "testing"

func TestVide(t *testing.T) {
	if !bienrange([]int{}) {
		t.Fail()
	}
}

func TestUn(t *testing.T) {
	if !bienrange([]int{1}) {
		t.Fail()
	}
}

func TestOk(t *testing.T) {
	if !bienrange([]int{1, 2, 5, 9, 12, 153, 1024}) {
		t.Fail()
	}
}

func TestPasOk(t *testing.T) {
	if bienrange([]int{1, 2, 5, 12, 9, 153, 1024}) {
		t.Fail()
	}
}
