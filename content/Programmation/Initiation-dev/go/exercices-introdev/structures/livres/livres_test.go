package livres

import "testing"

func TestVide(t *testing.T) {
	tab := []livre{}
	trier(tab, true)
	if len(tab) != 0 {
		t.Error("Après premier tri, la taille du tableau vide a changé")
		return
	}
	trier(tab, false)
	if len(tab) != 0 {
		t.Error("Après deuxième tri, la taille du tableau vide a changé")
	}
}

func TestPrix(t *testing.T) {
	l1 := livre{titre: "L'empire ultime", prix: 10.90, numPages: 928}
	l2 := livre{titre: "Elantris", prix: 11.20, numPages: 800}
	l3 := livre{titre: "Warbreaker", prix: 10.90, numPages: 984}
	l4 := livre{titre: "Axiomatique", prix: 8.90, numPages: 512}
	l5 := livre{titre: "Radieux", prix: 8.20, numPages: 480}
	l6 := livre{titre: "Océanique", prix: 9.90, numPages: 800}
	tab := []livre{l1, l2, l3, l4, l5, l6}
	trier(tab, true)
	if len(tab) != 6 {
		t.Error("Après tri, la taille du tableau a changé")
		return
	}
	if tab[0] != l5 || tab[1] != l4 || tab[2] != l6 || tab[3] != l1 || tab[4] != l3 || tab[5] != l2 {
		t.Error(disperror(l5, l4, l6, l1, l3, l2, tab))
	}
}

func TestPages(t *testing.T) {
	l1 := livre{titre: "L'empire ultime", prix: 10.90, numPages: 928}
	l2 := livre{titre: "Elantris", prix: 11.20, numPages: 800}
	l3 := livre{titre: "Warbreaker", prix: 10.90, numPages: 984}
	l4 := livre{titre: "Axiomatique", prix: 8.90, numPages: 512}
	l5 := livre{titre: "Radieux", prix: 8.20, numPages: 480}
	l6 := livre{titre: "Océanique", prix: 9.90, numPages: 800}
	tab := []livre{l1, l2, l3, l4, l5, l6}
	trier(tab, false)
	if len(tab) != 6 {
		t.Error("Après tri, la taille du tableau a changé")
		return
	}
	if tab[0] != l3 || tab[1] != l1 || tab[2] != l2 || tab[3] != l6 || tab[4] != l4 || tab[5] != l5 {
		t.Error(disperror(l3, l1, l2, l6, l4, l5, tab))
	}
}

// affichage de l'erreur
func disperror(l1, l2, l3, l4, l5, l6 livre, tab []livre) string {
	return "L'ordre attendu est " + l1.titre + ", " + l2.titre + ", " + l3.titre + ", " + l4.titre + ", " + l5.titre + ", " + l6.titre + " mais vous indiquez " + tab[0].titre + ", " + tab[1].titre + ", " + tab[2].titre + ", " + tab[3].titre + ", " + tab[4].titre + ", " + tab[5].titre
}
