package longueur

import "testing"

func TestVide(t *testing.T) {
	tab := []string{}
	trier(tab)
	if len(tab) != 0 {
		t.Error("Après tri, la taille du tableau vide a changé")
	}
}

func TestTab(t *testing.T) {
	tab := []string{"je", "souhaite", "un", "bon", "test", "a", "tout", "le", "monde"}
	trier(tab)
	if len(tab) != 9 {
		t.Error("Après tri, la taille du tableau a changé")
		return
	}
	if tab[0] != "souhaite" ||
		tab[1] != "monde" ||
		tab[2] != "test" ||
		tab[3] != "tout" ||
		tab[4] != "bon" ||
		tab[5] != "je" ||
		tab[6] != "le" ||
		tab[7] != "un" ||
		tab[8] != "a" {
		t.Error("Le résultat attendu est [souhaite monde test tout bon Je le un à], mais vous retournez", tab)
	}
}
