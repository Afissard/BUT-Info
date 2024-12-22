package tabtab

import (
	"testing"
)

func TestVide1(t *testing.T) {
	if _, err := compte(0, nil); err != errPasTabTab {
		t.Fail()
	}
}

func TestVide2(t *testing.T) {
	if num, err := compte(0, [][]int{}); num != 0 || err != nil {
		t.Fail()
	}
}

func TestVide3(t *testing.T) {
	if num, err := compte(0, [][]int{[]int{}, []int{}, []int{}}); num != 0 || err != nil {
		t.Fail()
	}
}

func TestCompte1(t *testing.T) {
	if num, err := compte(
		1, [][]int{[]int{1}, []int{2}, []int{3}, []int{1}},
	); num != 2 || err != nil {
		t.Fail()
	}
}

func TestCompte2(t *testing.T) {
	if num, err := compte(
		1, [][]int{[]int{1, 2, 3, 1}},
	); num != 2 || err != nil {
		t.Fail()
	}
}

func TestCompte3(t *testing.T) {
	if num, err := compte(
		1, [][]int{
			[]int{1, 2, 3, 1},
			[]int{2, 3, 1, 2},
			[]int{3, 1, 2, 3},
		},
	); num != 4 || err != nil {
		t.Fail()
	}
}

func TestCompte4(t *testing.T) {
	if num, err := compte(
		1, [][]int{
			[]int{1, 2, 3, 1},
			[]int{2, 3, 1},
			[]int{3, 1, 2, 3},
		},
	); num != 4 || err != nil {
		t.Fail()
	}
}

func TestCompte5(t *testing.T) {
	if num, err := compte(
		1, [][]int{
			[]int{1, 2, 3, 1},
			[]int{2, 3, 1},
			[]int{3, 1, 2, 3},
			[]int{},
		},
	); num != 4 || err != nil {
		t.Fail()
	}
}

func TestError1(t *testing.T) {
	if _, err := compte(
		1, [][]int{
			[]int{1, 2, 3, 1},
			nil,
			[]int{3, 1, 2, 3},
			[]int{},
		},
	); err != errPasTabTab {
		t.Fail()
	}
}

func TestError2(t *testing.T) {
	if _, err := compte(
		1, [][]int{
			[]int{},
			nil,
			[]int{},
			[]int{},
		},
	); err != errPasTabTab {
		t.Fail()
	}
}

func TestError3(t *testing.T) {
	if _, err := compte(
		1, [][]int{
			nil,
			nil,
			nil,
		},
	); err != errPasTabTab {
		t.Fail()
	}
}
