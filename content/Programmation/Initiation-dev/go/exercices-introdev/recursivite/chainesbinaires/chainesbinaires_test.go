package chainesbinaires

import (
	"sort"
	"testing"
)

func TestVide(t *testing.T) {
	if len(calculeChaines(-1)) != 0 {
		t.Fail()
	}
}

func TestZero(t *testing.T) {
	chaines := calculeChaines(0)
	if len(chaines) != 1 {
		t.Fail()
	} else if chaines[0] != "" {
		t.Fail()
	}
}

func TestUn(t *testing.T) {
	chaines := calculeChaines(1)
	if len(chaines) != 2 {
		t.Fail()
	} else if chaines[0] != "0" && chaines[0] != "1" {
		t.Fail()
	} else if chaines[1] != "0" && chaines[1] != "1" {
		t.Fail()
	} else if chaines[0] == chaines[1] {
		t.Fail()
	}
}

func TestTrois(t *testing.T) {
	if !memeContenu(
		calculeChaines(3),
		[]string{"000", "001", "010", "100", "101"},
	) {
		t.Fail()
	}
}

func TestCinq(t *testing.T) {
	if !memeContenu(
		calculeChaines(5),
		[]string{
			"00000", "00001", "00010", "00100",
			"00101", "01000", "01001", "01010",
			"10000", "10001", "10010", "10100",
			"10101",
		},
	) {
		t.Fail()
	}
}

// Fonction utilitaire pour les tests
func memeContenu(t1, t2 []string) bool {
	if len(t1) != len(t2) {
		return false
	}
	tt1 := make([]string, len(t1))
	copy(tt1, t1)
	tt2 := make([]string, len(t2))
	copy(tt2, t2)
	sort.Slice(tt1, func(i, j int) bool { return tt1[i] < tt1[j] })
	sort.Slice(tt2, func(i, j int) bool { return tt2[i] < tt2[j] })
	for i, v := range tt1 {
		if v != tt2[i] {
			return false
		}
	}
	return true
}
