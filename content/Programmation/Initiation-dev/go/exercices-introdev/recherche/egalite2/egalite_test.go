package egalite

import "testing"

func TestVide(t *testing.T) {
	if !egalite([]int{}, []int{}) {
		t.Fail()
	}
	if !egalite(nil, []int{}) {
		t.Fail()
	}
	if !egalite([]int{}, nil) {
		t.Fail()
	}
	if !egalite(nil, nil) {
		t.Fail()
	}
}

func TestVrai(t *testing.T) {
	if !egalite([]int{1, 2, 3, 4, 5}, []int{2, 4, 3, 1, 5}) {
		t.Fail()
	}
}

func TestFaux(t *testing.T) {
	if egalite([]int{1, 2, 3, 4, 5}, []int{1, 4, 3, 6, 5}) {
		t.Fail()
	}
}

func TestLongueur(t *testing.T) {
	if egalite([]int{}, []int{1}) {
		t.Fail()
	}
	if egalite([]int{1}, []int{}) {
		t.Fail()
	}
}

func TestMultipleVrai(t *testing.T) {
	if !egalite([]int{1, 2, 3, 4, 5, 1, 1, 2, 3, 3, 3}, []int{1, 2, 4, 2, 3, 3, 1, 3, 1, 3, 5}) {
		t.Fail()
	}
}

func TestMultipleFaux(t *testing.T) {
	if egalite([]int{1, 2, 2, 3, 4, 5}, []int{1, 4, 3, 6, 5}) {
		t.Fail()
	}
}
