package somme

import "testing"

func TestVide(t *testing.T) {
	res := somme(nil)
	if res != 0 {
		t.Error("somme(nil) devrait retourner 0 mais retourne", res)
	}
	res = somme([]int{})
	if res != 0 {
		t.Error("somme([]int{}) devrait retourner 0 mais retourne", res)
	}
}

func TestSimple(t *testing.T) {
	res := somme([]int{1, 1, 1, 1, 1})
	if res != 5 {
		t.Error("somme([]int{1, 1, 1, 1, 1}) devrait retourner 5",
			"mais retourne", res)
	}
}
