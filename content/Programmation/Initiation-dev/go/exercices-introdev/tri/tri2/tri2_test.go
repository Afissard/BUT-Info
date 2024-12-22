package tri2

import "testing"

func TestVide(t *testing.T) {
	if len(triabs([]int{})) != 0 {
		t.Error("triabs([]int{}) devrait retourner un tableau vide mais ce n'est pas le cas")
	}
	if len(triabs(nil)) != 0 {
		t.Error("triabs(nil) devrait retourner un tableau vide mais ce n'est pas le cas")
	}
}

func TestNonModif(t *testing.T) {
	tinit := []int{2, -3, -1, 4, 5}
	triabs(tinit)
	if len(tinit) != 5 {
		t.Error("La fonction triabs modifie la taille de son entrée")
	} else if tinit[0] != 2 || tinit[1] != -3 || tinit[2] != -1 || tinit[3] != 4 ||
		tinit[4] != 5 {
		t.Error("La fonction triabs modifie le contenu de son entrée")
	}
}

func TestTri(t *testing.T) {
	tinit := []int{2, -3, -1, 4, 5}
	tfin := triabs(tinit)
	if len(tfin) != 5 {
		t.Error("La fonction triabs retourne un tableau qui n'a pas la bonne taille")
	} else if tfin[0] != 5 || tfin[1] != 4 || tfin[2] != -3 || tfin[3] != 2 ||
		tfin[4] != -1 {
		t.Error("La fonction tri ne trie pas correctement son entrée")
	}
}

// Ajouté après le test machine 2021-2022
func TestTri2(t *testing.T) {
	tinit := []int{-1, -2, -3, -4, -5}
	tfin := triabs(tinit)
	if len(tfin) != 5 {
		t.Error("La fonction triabs retourne un tableau qui n'a pas la bonne taille")
	} else if tfin[0] != -5 || tfin[1] != -4 || tfin[2] != -3 || tfin[3] != -2 ||
		tfin[4] != -1 {
		t.Error("La fonction tri ne trie pas correctement son entrée")
	}
}
