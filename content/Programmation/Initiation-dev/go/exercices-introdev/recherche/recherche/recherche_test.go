package recherche

import "testing"

func TestVide(t *testing.T) {
	trouve, _ := recherche(nil, 0)
	if trouve {
		t.Error("recherche(nil, 0) retourne true mais devrait retourner false")
	}
	trouve, _ = recherche([]int{}, 0)
	if trouve {
		t.Error("recherche([]int{}, 0) retourne true mais devrait retourner false")
	}
}

func TestAbsent(t *testing.T) {
	trouve, _ := recherche([]int{1, 2, 3, 4, 5}, 6)
	if trouve {
		t.Error("recherche([]int{1, 2, 3, 4, 5}, 6) retourne true",
			"mais devrait retourner false")
	}
}

func TestPresent(t *testing.T) {
	trouve, pos := recherche([]int{1, 2, 3, 4, 5}, 4)
	if !trouve {
		t.Error("recherche([]int{1, 2, 3, 4, 5}, 4) retourne false",
			"mais devrait retourner true")
	} else if pos != 3 {
		t.Error("recherche([]int{1, 2, 3, 4, 5}, 4) retourne", pos,
			"mais devrait retourner 3")
	}
}
