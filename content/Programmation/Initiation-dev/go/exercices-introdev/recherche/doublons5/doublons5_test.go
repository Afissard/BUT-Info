package doublons5

import "testing"

func TestVide(t *testing.T) {
	if !doublons([]int{}) {
		t.Fail()
	}
}

func TestOk1(t *testing.T) {
	if !doublons([]int{1, 2, 3, 4, 5}) {
		t.Fail()
	}
}

func TestOk2(t *testing.T) {
	if !doublons([]int{2, 1, 5, 3, 4}) {
		t.Fail()
	}
}

func TestOk3(t *testing.T) {
	if !doublons([]int{2, 1, 5, 3, 4, -1, -2, -3}) {
		t.Fail()
	}
}

func TestPasOk1(t *testing.T) {
	if doublons([]int{1, 1}) {
		t.Fail()
	}
}

func TestPasOk2(t *testing.T) {
	if doublons([]int{5, 3, 2, 1, 4, 1, 2, 3, 4}) {
		t.Fail()
	}
}

func TestPasOk3(t *testing.T) {
	if doublons([]int{5, 3, -2, -1, 4, 1, -2, -3, -4}) {
		t.Fail()
	}
}
