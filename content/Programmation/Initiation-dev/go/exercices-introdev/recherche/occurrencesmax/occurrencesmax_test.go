package occurrencesmax

import (
	"testing"
)

func TestOcc1(t *testing.T) {
	if _, occ := occurrencesmax(nil); occ != 0 {
		t.Fail()
	}
}

func TestOcc2(t *testing.T) {
	if _, occ := occurrencesmax([]int{}); occ != 0 {
		t.Fail()
	}
}

func TestOcc3(t *testing.T) {
	if _, occ := occurrencesmax([]int{1, 2, 3, 4, 5, 6, 7, 8}); occ != 1 {
		t.Fail()
	}
}

func TestOcc4(t *testing.T) {
	if _, occ := occurrencesmax([]int{5, 5, 5, 5, 5}); occ != 5 {
		t.Fail()
	}
}

func TestOcc5(t *testing.T) {
	if _, occ := occurrencesmax([]int{8, 1, 3, 4, 1, 6, 5, 2, 4, 5, 3, 5, 2, 6, 7, 8, 7}); occ != 2 {
		t.Fail()
	}
}

func TestVal1(t *testing.T) {
	if n, _ := occurrencesmax([]int{5}); n != 5 {
		t.Fail()
	}
}

func TestVal2(t *testing.T) {
	if n, _ := occurrencesmax([]int{8, 1, 3, 4, 1, 6, 5, 2, 4, 5, 3, 5, 2, 6, 7, 8, 7}); n != 8 {
		t.Fail()
	}
}
