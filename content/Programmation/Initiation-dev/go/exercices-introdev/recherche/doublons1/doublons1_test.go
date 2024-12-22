package doublons1

import "testing"

func TestVide(t *testing.T) {
	if !doublons([]int{}) {
		t.Fail()
	}
}

func TestOk(t *testing.T) {
	if !doublons([]int{1, 2, 3, 4, 5}) {
		t.Fail()
	}
}

func TestPasOk1(t *testing.T) {
	if doublons([]int{7, 3, 1, 2, 9}) {
		t.Fail()
	}
}

func TestPasOk2(t *testing.T) {
	if doublons([]int{5, 3, 2, 1, 4}) {
		t.Fail()
	}
}

func TestPasOk3(t *testing.T) {
	if doublons([]int{1, 2, 3, 4, 4}) {
		t.Fail()
	}
}

func TestPasOk4(t *testing.T) {
	if doublons([]int{1, 2, 3, 4, 6}) {
		t.Fail()
	}
}
