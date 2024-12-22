package conway

import "testing"

func TestU0(t *testing.T) {
	tab := conway(0)
	if len(tab) != 1 || tab[0] != 1 {
		t.Error("conway(0) retourne ", tab, " mais devrait retourner [1]")
	}
}

func TestU3(t *testing.T) {
	tab := conway(3)
	if len(tab) != 4 || tab[0] != 1 || tab[1] != 2 || tab[2] != 1 || tab[3] != 1 {
		t.Error("conway(2) retourne ", tab, " mais devrait retourner [1 2 1 1]")
	}
}

func TestU5(t *testing.T) {
	tab := conway(5)
	if len(tab) != 6 || tab[0] != 3 || tab[1] != 1 || tab[2] != 2 || tab[3] != 2 ||
		tab[4] != 1 || tab[5] != 1 {
		t.Error("conway(2) retourne ", tab, " mais devrait retourner [3 1 2 2 1 1]")
	}
}
