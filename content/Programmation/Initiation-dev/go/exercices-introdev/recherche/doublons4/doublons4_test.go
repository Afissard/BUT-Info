package doublons4

import "testing"

func TestVide(t *testing.T) {
	_, ok := doublons([]int{})
	if !ok {
		t.Fail()
	}
}

func TestOk1(t *testing.T) {
	k, ok := doublons([]int{1, 2, 3, 4, 5})
	if !ok || k != 1 {
		t.Fail()
	}
}

func TestOk2(t *testing.T) {
	k, ok := doublons([]int{2, 1, 5, 3, 4})
	if !ok || k != 1 {
		t.Fail()
	}
}

func TestOk3(t *testing.T) {
	k, ok := doublons([]int{2, 2, 1, 3, 3, 3, 1, 1, 2})
	if !ok || k != 3 {
		t.Fail()
	}
}

func TestPasOk1(t *testing.T) {
	_, ok := doublons([]int{7, 3, 1, 2, 9})
	if ok {
		t.Fail()
	}
}

func TestPasOk2(t *testing.T) {
	_, ok := doublons([]int{5, 3, 2, 1, 4, 1, 2, 3, 4})
	if ok {
		t.Fail()
	}
}

func TestPasOk3(t *testing.T) {
	_, ok := doublons([]int{-1})
	if ok {
		t.Fail()
	}
}

func TestPasOk4(t *testing.T) {
	_, ok := doublons([]int{1, 2, 3, 4, 4, 4, 4, 4})
	if ok {
		t.Fail()
	}
}
