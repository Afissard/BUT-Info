package doublons2

import "testing"

func TestVide(t *testing.T) {
	if !doublons(nil) {
		t.Error("doublons(nil) devrait retourner true")
	}
	if !doublons([]int{}) {
		t.Error("doublons([]int{}) devrait retourner true")
	}
}

func TestDouble(t *testing.T) {
	if doublons([]int{1, 1}) {
		t.Error("doublons([]int{1, 1}) devrait retourner false")
	}
	if doublons([]int{1, 2, 3, 4, 5, 1}) {
		t.Error("doublons([]int{1, 2, 3, 4, 5, 1}) devrait retourner false")
	}
}

func TestOk(t *testing.T) {
	if !doublons([]int{1, 2, 3, 4, 5, 6}) {
		t.Error("doublons([]int{1, 2, 3, 4, 5, 6}) devrait retourner true")
	}
}

func TestDehors(t *testing.T) {
	if doublons([]int{2, 3, 4, 5, 6, 7}) {
		t.Error("doublons([]int{2, 3, 4, 5, 6, 7}) devrait retourner false")
	}
}

// Ajoutés après le premier test machine 2021-2022

func TestOk2(t *testing.T) {
	// Le tableau n'est pas nécessairement trié
	if !doublons([]int{5, 3, 1, 6, 2, 4}) {
		t.Fail()
	}
}

func TestSomme(t *testing.T) {
	// Quelques étudiants comparaient 1 + 2 + ... + n avec la somme des nombres
	// du tableau, ceci n'est pas correct, comme le montre ce cas de test
	if doublons([]int{1, 3, 4, 3, 4, 6}) {
		t.Fail()
	}
}

func TestNegatif(t *testing.T) {
	// Quelques étudiants comptaient le nombre d'apparitions de chaque
	// valeur dans le tableau, sans vérifier que ces valeurs étaient bien
	// entre 1 et len(t)
	if doublons([]int{-1}) {
		t.Fail()
	}
}

func TestTropGrand(t *testing.T) {
	// Quelques étudiants comptaient le nombre d'apparitions de chaque
	// valeur dans le tableau, sans vérifier que ces valeurs étaient bien
	// entre 1 et len(t)
	if doublons([]int{2}) {
		t.Fail()
	}
}
