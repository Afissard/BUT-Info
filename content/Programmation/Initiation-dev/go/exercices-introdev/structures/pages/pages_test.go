package pages

import "testing"

func TestVide(t *testing.T) {
	tab := []livre{}
	_, trouve := trouver(tab, 250)
	if trouve {
		t.Error("Vous réussissez à trouver un livre dans un tableau vide… étrange")
	}
}

func TestExiste(t *testing.T) {
	l1 := livre{titre: "L'empire ultime", prix: 10.90, numPages: 928}
	l2 := livre{titre: "Elantris", prix: 11.20, numPages: 800}
	l3 := livre{titre: "Warbreaker", prix: 10.90, numPages: 984}
	l4 := livre{titre: "Axiomatique", prix: 8.90, numPages: 512}
	l5 := livre{titre: "Radieux", prix: 8.20, numPages: 480}
	l6 := livre{titre: "Océanique", prix: 9.90, numPages: 800}
	tab := []livre{l1, l2, l3, l4, l5, l6}
	l, trouve := trouver(tab, 512)
	if !trouve {
		t.Error("Vous n'avez pas trouvé de livre de 512 pages dans le tableau, pourtant il existe")
		return
	}
	if l != l4 {
		t.Error("Vous avez trouvé que", l.titre, "fait 512 pages mais il fallait trouver", l4.titre)
	}
}

func TestPasExiste(t *testing.T) {
	l1 := livre{titre: "L'empire ultime", prix: 10.90, numPages: 928}
	l2 := livre{titre: "Elantris", prix: 11.20, numPages: 800}
	l3 := livre{titre: "Warbreaker", prix: 10.90, numPages: 984}
	l4 := livre{titre: "Axiomatique", prix: 8.90, numPages: 512}
	l5 := livre{titre: "Radieux", prix: 8.20, numPages: 480}
	l6 := livre{titre: "Océanique", prix: 9.90, numPages: 800}
	tab := []livre{l1, l2, l3, l4, l5, l6}
	l, trouve := trouver(tab, 801)
	if trouve {
		t.Error("Non, il n'y a pas de livre de 801 pages dans le tableau, pourtant vous trouvez", l.titre)
	}
}
