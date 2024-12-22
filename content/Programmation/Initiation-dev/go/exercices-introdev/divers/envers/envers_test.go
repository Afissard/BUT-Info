package envers

import "testing"

func TestVide(t *testing.T) {
	tab := []int{}
	envers(tab)
	if len(tab) != 0 {
		t.Error("Votre fonction donne ", tab, " au lieu du tableau vide")
	}
}

func Test1(t *testing.T) {
	tab := []int{1, 2, 3}
	envers(tab)
	if len(tab) != 3 {
		t.Error("Votre fonction modifie la longueur des tableaux")
	} else if tab[0] != 3 || tab[1] != 2 || tab[2] != 1 {
		t.Error("Votre fonction donne ", tab, " au lieu de [3 2 1]")
	}
}

func Test2(t *testing.T) {
	tab := []int{1, 2, 3, 4}
	envers(tab)
	if len(tab) != 4 {
		t.Error("Votre fonction modifie la longueur des tableaux")
	} else if tab[0] != 4 || tab[1] != 3 || tab[2] != 2 || tab[3] != 1 {
		t.Error("Votre fonction donne ", tab, " au lieu de [4 3 2 1]")
	}
}
