package multipleDeTrois

import "testing"

func Test1(t *testing.T) {
	tab := []int{1, 3, 5, 7}
	res := unMultipleDeTrois(tab)
	if res != 3 {
		t.Error("Dans", tab, "3 est le seul multiple de trois mais vous retournez", res)
	}
}

func Test2(t *testing.T) {
	tab := []int{1, 5, 7, 8}
	res := unMultipleDeTrois(tab)
	if res != -1 {
		t.Error("Dans", tab, "il n'y a pas de multiple de 3 et pourtant vous retournez", res)
	}
}

// Ajouté après le test 1 de 2022-2023

func Test3(t *testing.T) {
	tab := []int{11, 7, 5, 4, 12, 8, 1, 5, 7, 8}
	res := unMultipleDeTrois(tab)
	if res != 12 {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	tab := []int{7, 5, 4, 2, 1, 5, 7, 8}
	res := unMultipleDeTrois(tab)
	if res != -1 {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	tab := []int{}
	res := unMultipleDeTrois(tab)
	if res != -1 {
		t.Fail()
	}
}
