package vraifaux

import "testing"

func TestVide(t *testing.T) {
	tab := []bool{}
	trier(tab)
	if len(tab) != 0 {
		t.Error("Après tri, la taille du tableau vide a changé")
	}
}

func TestTab(t *testing.T) {
	tab := []bool{true, true, false, false, true, false, false, true}
	trier(tab)
	if len(tab) != 8 {
		t.Error("Après tri, la taille du tableau a changé")
		return
	}
	if !tab[0] || !tab[1] || !tab[2] || !tab[3] || tab[4] || tab[5] || tab[6] || tab[7] {
		t.Error("Le résultat attendu est [true true true true false false false false], mais vous retournez", tab)
	}
}

// Tests ajoutés après le test machine

func TestTab1(t *testing.T) {
	tab := []bool{true, true, true}
	trier(tab)
	if len(tab) != 3 {
		t.Error("Après tri, la taille du tableau a changé")
		return
	}
	if !tab[0] || !tab[1] || !tab[2] {
		t.Error("Le résultat attendu est [true true true], mais vous retournez", tab)
	}
}

func TestTab2(t *testing.T) {
	tab := []bool{false, false, false}
	trier(tab)
	if len(tab) != 3 {
		t.Error("Après tri, la taille du tableau a changé")
		return
	}
	if tab[0] || tab[1] || tab[2] {
		t.Error("Le résultat attendu est [false false false], mais vous retournez", tab)
	}
}

func TestTab3(t *testing.T) {
	tab := []bool{true, true, false, false}
	trier(tab)
	if len(tab) != 4 {
		t.Error("Après tri, la taille du tableau a changé")
		return
	}
	if !tab[0] || !tab[1] || tab[2] || tab[3] {
		t.Error("Le résultat attendu est [true true false false], mais vous retournez", tab)
	}
}

func TestTab4(t *testing.T) {
	tab := []bool{false, false, true, true}
	trier(tab)
	if len(tab) != 4 {
		t.Error("Après tri, la taille du tableau a changé")
		return
	}
	if !tab[0] || !tab[1] || tab[2] || tab[3] {
		t.Error("Le résultat attendu est [true true false false], mais vous retournez", tab)
	}
}
