package moitie

import "testing"

func Test1(t *testing.T) {
	tab := []int{1, 1, 1, -1, -1, -1}
	if !moitieDePositifs(tab) {
		t.Error("Le tableau", tab, "a au moins la moitié des nombres qu'il contient qui sont strictement supérieurs à 0, pourtant vous affirmez le contraire")
	}
}

func Test2(t *testing.T) {
	tab := []int{1, 1, -1, -1, -1}
	if moitieDePositifs(tab) {
		t.Error("Le tableau", tab, "a moins de la moitié des nombres qu'il contient qui sont strictement supérieurs à 0, pourtant vous affirmez le contraire")
	}
}

// Ajouté après le test 1 de 2022-2023

func Test3(t *testing.T) {
	tab := []int{}
	if !moitieDePositifs(tab) {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	tab := []int{5, 6, 7, -2, 6, -3, -4}
	if !moitieDePositifs(tab) {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	tab := []int{-5, -6, -7, -2, 6, 3, 4, 1, -2, 3}
	if !moitieDePositifs(tab) {
		t.Fail()
	}
}
