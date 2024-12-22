package sousensembles

import (
	"sort"
	"testing"
)

func TestErrorVide(t *testing.T) {
	if _, err := sousEnsembles(nil); err != errPasEnsemble {
		t.Fail()
	}
}

func TestErrorPasEnsemble(t *testing.T) {
	if _, err := sousEnsembles([]int{1, 2, 1}); err != errPasEnsemble {
		t.Fail()
	}
}

func TestErrorPasEnsemble2(t *testing.T) {
	if _, err := sousEnsembles([]int{1, 2, 3, 2}); err != errPasEnsemble {
		t.Fail()
	}
}

func TestSingleton(t *testing.T) {
	pe, err := sousEnsembles([]int{})
	if err != nil {
		t.Fail()
	} else if len(pe) != 1 {
		t.Fail()
	} else if len(pe[0]) != 0 {
		t.Fail()
	} else if pe[0] == nil {
		t.Fail()
	}
}

func Test1(t *testing.T) {
	pe, err := sousEnsembles([]int{0})
	if err != nil {
		t.Fail()
	} else if len(pe) != 2 {
		t.Fail()
	} else if pe[0] == nil || pe[1] == nil {
		t.Fail()
	} else if len(pe[0]) != 0 && len(pe[0]) != 1 {
		t.Fail()
	} else if len(pe[1]) != 0 && len(pe[1]) != 1 {
		t.Fail()
	} else if len(pe[0]) == 1 && pe[0][0] != 0 {
		t.Fail()
	} else if len(pe[1]) == 1 && pe[1][0] != 0 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	pe, err := sousEnsembles([]int{1, 2})
	if err != nil {
		t.Fail()
	} else if !memeContenu(pe, [][]int{{}, {1}, {2}, {1, 2}}) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	pe, err := sousEnsembles([]int{1, 2, 3})
	if err != nil {
		t.Fail()
	} else if !memeContenu(
		pe,
		[][]int{
			{}, {1}, {2}, {3},
			{1, 2}, {1, 3}, {2, 3},
			{1, 2, 3},
		},
	) {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	pe, err := sousEnsembles([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Fail()
	} else if !memeContenu(
		pe,
		[][]int{
			{},
			{1}, {2}, {3}, {4}, {5},
			{1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 3}, {2, 4}, {2, 5}, {3, 4}, {3, 5}, {4, 5},
			{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5},
			{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5},
			{1, 2, 3, 4, 5},
		},
	) {
		t.Fail()
	}
}

// Fonction utilitaire pour les tests
func memeContenu(t1, t2 [][]int) bool {
	if len(t1) != len(t2) {
		return false
	}
	tt1 := make([][]int, len(t1))
	tt2 := make([][]int, len(t2))
	for i := 0; i < len(t1); i++ {
		tt1[i] = make([]int, len(t1[i]))
		copy(tt1[i], t1[i])
		trieTout(tt1)
		tt2[i] = make([]int, len(t2[i]))
		copy(tt2[i], t2[i])
		trieTout(tt2)
	}
	sort.Slice(tt1, func(i, j int) bool { return estAvant(tt1[i], tt1[j]) })
	sort.Slice(tt2, func(i, j int) bool { return estAvant(tt2[i], tt2[j]) })
	for i, v := range tt1 {
		if !egal(v, tt2[i]) {
			return false
		}
	}
	return true
}

func trieTout(t [][]int) {
	for _, tt := range t {
		sort.Slice(tt, func(i, j int) bool { return tt[i] < tt[j] })
	}
}

func estAvant(t1, t2 []int) bool {
	if len(t1) < len(t2) {
		return true
	}
	if len(t1) > len(t2) {
		return false
	}
	for i := 0; i < len(t1); i++ {
		if t1[i] < t2[i] {
			return true
		}
		if t1[i] > t2[i] {
			return false
		}
	}
	return true
}

func egal(t1, t2 []int) bool {
	if len(t1) != len(t2) {
		return false
	}
	for i := 0; i < len(t1); i++ {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}
