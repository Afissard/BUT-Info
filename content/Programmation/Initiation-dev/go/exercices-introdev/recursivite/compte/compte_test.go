package compte

import "testing"

func TestVide(t *testing.T) {
	num := compte(nil, 0)
	if num != 0 {
		t.Error("compte(nil, 0) retourne", num, "mais devrait retourner 0")
	}
	num = compte([]int{}, 0)
	if num != 0 {
		t.Error("compte([]int{}, 0) retourne", num, "mais devrait retourner 0")
	}
}

func TestAbsent(t *testing.T) {
	num := compte([]int{1, 2, 3, 4, 5}, 6)
	if num != 0 {
		t.Error("compte([]int{1, 2, 3, 4, 5}, 6) retourne", num,
			"mais devrait retourner 0")
	}
}

func TestPresent(t *testing.T) {
	num := compte([]int{1, 2, 1, 2, 1, 2}, 2)
	if num != 3 {
		t.Error("compte([]int{1, 2, 1, 2, 1, 2}, 2) retourne", num,
			"mais devrait retourner 3")
	}
}
