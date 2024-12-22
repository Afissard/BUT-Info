package pairs

import "testing"

func TestVide(t *testing.T) {
	num := pairs([]int{})
	if num != 0 {
		t.Error("Le nombre de nombres pairs dans un tableau vide est 0 mais vous en trouvez", num)
	}
}

func TestExemple(t *testing.T) {
	tab := []int{1, 3, 4, 2, 6, 7, 8, 5, 2, 4}
	num := pairs([]int{1, 3, 4, 2, 6, 7, 8, 5, 2, 4})
	if num != 6 {
		t.Error("Le nombre de nombres pairs dans le tableau", tab, "est 6 mais vous en trouvez", num)
	}
}
