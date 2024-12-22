package maxoccurrences

import (
	"testing"
)

func TestOcc1(t *testing.T) {
	if _, occ := maxoccurrences(nil); occ != 0 {
		t.Fail()
	}
}

func TestOcc2(t *testing.T) {
	if _, occ := maxoccurrences([]int{}); occ != 0 {
		t.Fail()
	}
}

func TestOcc3(t *testing.T) {
	if _, occ := maxoccurrences([]int{1, 2, 3, 4, 5, 6, 7, 8}); occ != 1 {
		t.Fail()
	}
}

func TestOcc4(t *testing.T) {
	if _, occ := maxoccurrences([]int{5, 5, 5, 5, 5}); occ != 5 {
		t.Fail()
	}
}

func TestOcc5(t *testing.T) {
	if _, occ := maxoccurrences([]int{1, 5, 1, 1, 5, 5, 1, 5, 5}); occ != 5 {
		t.Fail()
	}
}

func TestOcc6(t *testing.T) {
	if _, occ := maxoccurrences([]int{8, 1, 3, 4, 1, 6, 5, 2, 4, 5, 3, 5, 2, 6, 7, 8, 7}); occ != 3 {
		t.Fail()
	}
}

func TestVal1(t *testing.T) {
	if n, _ := maxoccurrences([]int{5}); n != 5 {
		t.Fail()
	}
}

func TestVal2(t *testing.T) {
	if n, _ := maxoccurrences([]int{8, 1, 3, 4, 1, 6, 5, 2, 4, 5, 3, 5, 2, 6, 7, 8, 7}); n != 5 {
		t.Fail()
	}
}

func TestVal3(t *testing.T) {
	if n, _ := maxoccurrences([]int{0, 1, 2, 3, 2, 1, 0, 1, 2}); n != 1 && n != 2 {
		t.Fail()
	}
}
