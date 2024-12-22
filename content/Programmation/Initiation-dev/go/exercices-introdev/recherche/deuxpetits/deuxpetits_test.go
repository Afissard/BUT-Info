package deuxpetits

import "testing"

func Test1(t *testing.T) {
	tab := []int{1, 2}
	res1, res2 := lesDeuxPlusPetits(tab)
	if res2 != 2 || res1 != 1 {
		t.Error("Pour la tableau [1 2] il faut retourner 1 comme plus petit (vous retournez", res1, "à la place) et 2 comme deuxième plus petit (vous retournez", res2, "à la place)")
	}
}

func Test2(t *testing.T) {
	tab := []int{33, 15, 2, 8, 1}
	res1, res2 := lesDeuxPlusPetits(tab)
	if res2 != 2 || res1 != 1 {
		t.Error("Pour la tableau [33 15 2 8 1] il faut retourner 1 comme plus petit (vous retournez", res1, "à la place) et 2 comme deuxième plus petit (vous retournez", res2, "à la place)")
	}
}

// ajoutés après le test machine

func Test3(t *testing.T) {
	tab := []int{7, 7}
	res1, res2 := lesDeuxPlusPetits(tab)
	if res2 != 7 || res1 != 7 {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	tab := []int{12, 11, 10, 9, 8, 7, 7, 8, 9, 10, 11}
	res1, res2 := lesDeuxPlusPetits(tab)
	if res2 != 7 || res1 != 7 {
		t.Fail()
	}
}
