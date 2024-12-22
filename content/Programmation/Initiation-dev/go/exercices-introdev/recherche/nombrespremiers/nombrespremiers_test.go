package nombrespremiers

import (
	"testing"
)

func TestVide1(t *testing.T) {
	if selectionPremiers(nil) != nil {
		t.Fail()
	}
}

func TestVide2(t *testing.T) {
	p := selectionPremiers([]int{})
	if p == nil || len(p) != 0 {
		t.Fail()
	}
}

func TestPremier(t *testing.T) {
	p := selectionPremiers([]int{2})
	if len(p) != 1 {
		t.Fail()
	} else if p[0] != 2 {
		t.Fail()
	}
}

func TestNonPremier(t *testing.T) {
	p := selectionPremiers([]int{4})
	if p == nil || len(p) != 0 {
		t.Fail()
	}
}

func TestUn(t *testing.T) {
	p := selectionPremiers([]int{1})
	if p == nil || len(p) != 0 {
		t.Fail()
	}
}

func TestZero(t *testing.T) {
	p := selectionPremiers([]int{0})
	if p == nil || len(p) != 0 {
		t.Fail()
	}
}

func TestNegatif(t *testing.T) {
	p := selectionPremiers([]int{-1})
	if p == nil || len(p) != 0 {
		t.Fail()
	}
}

func TestDouble(t *testing.T) {
	p := selectionPremiers([]int{2, 2})
	if len(p) != 2 {
		t.Fail()
	} else if p[0] != 2 || p[1] != 2 {
		t.Fail()
	}
}

func TestPlusieurs(t *testing.T) {
	p := selectionPremiers([]int{
		0, -1, -2, 12, 54,
		1, 2, 7, 7, 7,
		3, 3, 4, 4, 5,
		6, 7, 8, 9, 10,
		127, 12, 12, 12, 12,
	})
	if len(p) != 9 {
		t.Fail()
	}
}
