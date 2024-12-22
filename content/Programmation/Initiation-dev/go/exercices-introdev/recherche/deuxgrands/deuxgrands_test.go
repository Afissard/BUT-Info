package deuxGrands

import "testing"

func Test1(t *testing.T) {
	tab := []int{1, 2}
	res1, res2 := lesDeuxPlusGrands(tab)
	if res2 != 1 || res1 != 2 {
		t.Error("Les deux plus grands éléments du tableau", tab, "sont 1 et 2 mais vous retournez", res1, "(plus grand) et", res2, " (deuxième plus grand)")
	}
}

func Test2(t *testing.T) {
	tab := []int{1, 2, 3, 4, 5}
	res1, res2 := lesDeuxPlusGrands(tab)
	if res2 != 4 || res1 != 5 {
		t.Error("Les deux plus grands éléments du tableau", tab, "sont 4 et 5 mais vous retournez", res1, "(plus grand) et", res2, " (deuxième plus grand)")
	}
}

func Test3(t *testing.T) {
	tab := []int{5, 4, 3, 2, 1}
	res1, res2 := lesDeuxPlusGrands(tab)
	if res2 != 4 || res1 != 5 {
		t.Error("Les deux plus grands éléments du tableau", tab, "sont 4 et 5 mais vous retournez", res1, "(plus grand) et", res2, " (deuxième plus grand)")
	}
}

func Test4(t *testing.T) {
	tab := []int{2, 2, 1}
	res1, res2 := lesDeuxPlusGrands(tab)
	if res2 != 2 || res1 != 2 {
		t.Error("Les deux plus grands éléments du tableau", tab, "sont 2 et 2 mais vous retournez", res1, "(plus grand) et", res2, " (deuxième plus grand)")
	}
}

// Ajouté après le test 1 de 2022-2023

func Test5(t *testing.T) {
	tab := []int{33, 15, 2, 8, 125, 3, 654, 4, 59, 2, 2, 1}
	res1, res2 := lesDeuxPlusGrands(tab)
	if res2 != 125 || res1 != 654 {
		t.Fail()
	}
}

func Test6(t *testing.T) {
	tab := []int{33, 15, 654, 8, 125, 3, 654, 4, 59, 2, 2, 1}
	res1, res2 := lesDeuxPlusGrands(tab)
	if res2 != 654 || res1 != 654 {
		t.Fail()
	}
}
