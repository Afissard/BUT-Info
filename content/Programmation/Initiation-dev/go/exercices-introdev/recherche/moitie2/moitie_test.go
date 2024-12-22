package moitie

import "testing"

func Test1(t *testing.T) {
	tab := []int{1, 1, 1, 2, 2, 2}
	if !moitieDePairs(tab) {
		t.Error("Dans [1 1 1 2 2 2] la moitié des nombres sont pairs")
	}
}

func Test2(t *testing.T) {
	tab := []int{2, 2, 1, 1, 1}
	if moitieDePairs(tab) {
		t.Error("Dans [2 2 1 1 1] moins de la moitié des nombres sont pairs")
	}
}

func Test3(t *testing.T) {
	tab := []int{}
	if !moitieDePairs(tab) {
		t.Error("Dans le tableau vide la moitié des nombres sont pairs")
	}
}

// ajoutés après le test machine

func Test4(t *testing.T) {
	tab := []int{8, 10, 12, 3, 5, 98, 122, 127, 1, 5, 1024}
	if !moitieDePairs(tab) {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	tab := []int{8, 10, 11, 3, 5, 97, 123, 127, 1, 5, 1024, 3}
	if moitieDePairs(tab) {
		t.Fail()
	}
}
