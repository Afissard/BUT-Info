package doublons3

import "testing"

func TestVide1(t *testing.T) {
	if !doublons([]int{}, 0) {
		t.Fail()
	}
}

func TestVide2(t *testing.T) {
	if !doublons([]int{}, 5) {
		t.Fail()
	}
}

func TestOk1(t *testing.T) {
	if !doublons([]int{1, 2, 3, 4, 5}, 1) {
		t.Fail()
	}
}

func TestOk2(t *testing.T) {
	if !doublons([]int{2, 1, 5, 3, 4}, 1) {
		t.Fail()
	}
}

func TestOk3(t *testing.T) {
	if !doublons([]int{2, 2, 1, 3, 3, 3, 1, 1, 2}, 3) {
		t.Fail()
	}
}

func TestPasOk1(t *testing.T) {
	if doublons([]int{7, 3, 1, 2, 9}, 1) {
		t.Fail()
	}
}

func TestPasOk2(t *testing.T) {
	if doublons([]int{5, 3, 2, 1, 4}, 2) {
		t.Fail()
	}
}

func TestPasOk3(t *testing.T) {
	if doublons([]int{2, 2, 1, 3, 3, 3, 1, 1, 2, 1, 2, 3}, 3) {
		t.Fail()
	}
}

func TestPasOk4(t *testing.T) {
	if doublons([]int{2, 2, 1, 3, 3, 3, 1, 1, 2, 1, 2, 3}, 2) {
		t.Fail()
	}
}
