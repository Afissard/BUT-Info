package interval

import "testing"

func TestExiste(t *testing.T) {
	tab := []int{1, 3, 6, 12, 2}
	res := dansInterval(tab)
	if res != 6 {
		t.Error("Le tableau [1 3 6 12 2] ne contient que le nombre 6 qui convient, mais vous retournez", res)
	}
}

func TestNonExiste(t *testing.T) {
	tab := []int{1, 12, 3, 81}
	res := dansInterval(tab)
	if res != -1 {
		t.Error("Le tableau [1 12 3 81] ne contient aucun nombre qui convient, il faut retourner -1, mais vous vous retournez", res)
	}
}

// ajoutés après le test machine

func TestExiste2(t *testing.T) {
	tab := []int{127, 9, 1, 3, 7, 12, 2, 34, 128}
	res := dansInterval(tab)
	if res != 7 {
		t.Fail()
	}
}

func TestNonExiste2(t *testing.T) {
	tab := []int{4, 9, 1, 12, 3, 81, 10, 24, 57, 9, 9, 11}
	res := dansInterval(tab)
	if res != -1 {
		t.Fail()
	}
}

func TestVide(t *testing.T) {
	tab := []int{}
	res := dansInterval(tab)
	if res != -1 {
		t.Fail()
	}
}
