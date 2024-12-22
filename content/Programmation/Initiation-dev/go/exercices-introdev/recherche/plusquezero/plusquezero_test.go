package plusQueZero

import "testing"

func Test1(t *testing.T) {
	tab := []int{1, 1, 1}
	if !tousPositifs(tab) {
		t.Error("Le tableau", tab, "ne contient que des nombres qui sont strictement supérieurs à 0, pourtant vous affirmez le contraire")
	}
}

func Test2(t *testing.T) {
	tab := []int{1, 1, -1}
	if tousPositifs(tab) {
		t.Error("Le tableau", tab, "contient au moins un nombre qui n'est pas strictement supérieur à 0, pourtant vous affirmez le contraire")
	}
}

// Ajouté après le test 1 de 2022-2023

func Test3(t *testing.T) {
	tab := []int{}
	if !tousPositifs(tab) {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	tab := []int{1, 2, 3, 4, 5, 6, 5, 4, 2, 1, 12, -3, 2, 6, 5, 4, 7, 11}
	if tousPositifs(tab) {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	tab := []int{1, 2, 3, 4, 5, 6, 5, 4, 2, 1, 12, 3, 2, 6, 5, 4, 7, 11}
	if !tousPositifs(tab) {
		t.Fail()
	}
}
