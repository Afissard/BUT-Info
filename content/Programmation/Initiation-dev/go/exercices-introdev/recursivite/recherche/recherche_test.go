package recherche

import "testing"

func TestVide(t *testing.T) {
	trouve := recherche(nil, 0)
	if trouve {
		t.Error("recherche(nil, 0) retourne true mais devrait retourner false")
	}
	trouve = recherche([]int{}, 0)
	if trouve {
		t.Error("recherche([]int{}, 0) retourne true mais devrait retourner false")
	}
}

func TestAbsent(t *testing.T) {
	trouve := recherche([]int{1, 2, 3, 4, 5}, 6)
	if trouve {
		t.Error("recherche([]int{1, 2, 3, 4, 5}, 6) retourne true",
			"mais devrait retourner false")
	}
}

func TestPresent(t *testing.T) {
	trouve := recherche([]int{1, 2, 3, 4, 5}, 4)
	if !trouve {
		t.Error("recherche([]int{1, 2, 3, 4, 5}, 4) retourne false",
			"mais devrait retourner true")
	}
}

// Ajouté après le troisième test machine 2021-2022

func TestZero(t *testing.T) {
	if recherche([]int{1, 2, 3, 4, 5}, 0) {
		t.Fail()
	}
}
