package nombrespremiers

import (
	"sort"
	"testing"
)

func Test1(t *testing.T) {
	p := premiers(0)
	if p == nil || len(p) != 0 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	p := premiers(-5)
	if p == nil || len(p) != 0 {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	p := premiers(10)
	if !memeContenu(p, []int{2, 3, 5, 7}) {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	p := premiers(50)
	if !memeContenu(p, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}) {
		t.Fail()
	}
}

// Fonction utilitaire pour les tests
func memeContenu(t1, t2 []int) bool {
	if len(t1) != len(t2) {
		return false
	}
	sort.Slice(t1, func(i, j int) bool { return t1[i] < t1[j] })
	sort.Slice(t2, func(i, j int) bool { return t2[i] < t2[j] })
	for i := 0; i < len(t1); i++ {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}
