package valabs

import "testing"
import "fmt"

func TestVide(t *testing.T) {
	tab := []int{}
	valabs(tab)
	if len(tab) != 0 || tab == nil {
		t.Fail()
	}
	tab = nil
	valabs(tab)
	if tab != nil {
		t.Fail()
	}
}

func TestDecroissant(t *testing.T) {
	tab := []int{5, 4, -3, 2, -1}
	valabs(tab)
	if len(tab) != 5 {
		t.Fail()
	} else if tab[4] != 5 ||
		tab[3] != 4 ||
		tab[2] != -3 ||
		tab[1] != 2 ||
		tab[0] != -1 {
		t.Fail()
	}
}

func TestDejaTrie(t *testing.T) {
	tab := []int{-1, 2, -3, 4, 5}
	valabs(tab)
	if len(tab) != 5 {
		t.Fail()
	} else if tab[4] != 5 ||
		tab[3] != 4 ||
		tab[2] != -3 ||
		tab[1] != 2 ||
		tab[0] != -1 {
		t.Fail()
	}
}

func TestMelange(t *testing.T) {
	tab := []int{-3, 5, 2, -1, 4}
	valabs(tab)
	if len(tab) != 5 {
		t.Fail()
	} else if tab[4] != 5 ||
		tab[3] != 4 ||
		tab[2] != -3 ||
		tab[1] != 2 ||
		tab[0] != -1 {
		t.Fail()
	}
}

func TestEgaux(t *testing.T) {
	tab := []int{-3, 5, 3, -1, 4}
	valabs(tab)
	if len(tab) != 5 {
		t.Fail()
	} else if tab[4] != 5 ||
		tab[3] != 4 ||
		tab[2] != 3 ||
		tab[1] != -3 ||
		tab[0] != -1 {
		fmt.Print(tab)
		t.Fail()
	}
}
