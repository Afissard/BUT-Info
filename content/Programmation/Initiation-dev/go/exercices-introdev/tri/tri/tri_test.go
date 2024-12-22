package tri

import "testing"

func TestVide(t *testing.T) {
	if len(tri([]int{})) != 0 {
		t.Error("tri([]int{}) devrait retourner un tableau vide mais ce n'est pas le cas")
	}
	if len(tri(nil)) != 0 {
		t.Error("tri(nil) devrait retourner un tableau vide mais ce n'est pas le cas")
	}
}

func TestNonModif(t *testing.T) {
	tinit := []int{2, 3, 1, 4, 5}
	tri(tinit)
	if len(tinit) != 5 {
		t.Error("La fonction tri modifie la taille de son entrée")
	} else if tinit[0] != 2 || tinit[1] != 3 || tinit[2] != 1 || tinit[3] != 4 ||
		tinit[4] != 5 {
		t.Error("La fonction tri modifie le contenu de son entrée")
	}
}

func TestTri(t *testing.T) {
	tinit := []int{2, 3, 1, 4, 5}
	tfin := tri(tinit)
	if len(tfin) != 5 {
		t.Error("La fonction tri retourne un tableau qui n'a pas la bonne taille")
	} else if tfin[0] != 1 || tfin[1] != 2 || tfin[2] != 3 || tfin[3] != 4 ||
		tfin[4] != 5 {
		t.Error("La fonction tri ne trie pas correctement son entrée")
	}
}

// Ajouté après le test machine 2021-2022
func TestTri2(t *testing.T) {
	tinit := []int{2, -2, -1, 1, 1}
	tfin := tri(tinit)
	if len(tfin) != 5 {
		t.Error("La fonction tri retourne un tableau qui n'a pas la bonne taille")
	} else if tfin[0] != -2 || tfin[1] != -1 || tfin[2] != 1 || tfin[3] != 1 ||
		tfin[4] != 2 {
		t.Error("La fonction tri ne trie pas correctement son entrée")
	}
}
