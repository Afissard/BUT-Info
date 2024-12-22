package addptr

import "testing"

func Test(t *testing.T) {
	var x, y int = 1, 2
	var s int = addptr(&x, &y)
	if s != 3 {
		t.Error("add ne calcule pas la bonne somme, on devrait avoir 3 et on a ", s)
	}
	if x != 1 {
		t.Error("add modifie son premier argument, ça ne devrait pas être le cas")
	}
	if y != 2 {
		t.Error("add modifie son deuxième argument, ça ne devrait pas être le cas")
	}
}
