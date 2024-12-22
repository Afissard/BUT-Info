package varswitch

import "testing"

func TestSame(t *testing.T) {
	x, y := 1, 1
	varswitch(&x, &y)
	if x != 1 || y != 1 {
		t.Error("varswitch(&x, &y) avec x = 1 et y = 1",
			"devrait donner x = 1 et y = 1",
			"mais donne x =", x, "et y =", y)
	}
}

func TestDifferent(t *testing.T) {
	x, y := 1, 2
	varswitch(&x, &y)
	if x != 2 || y != 1 {
		t.Error("varswitch(&x, &y) avec x = 1 et y = 2",
			"devrait donner x = 2 et y = 1",
			"mais donne x =", x, "et y =", y)
	}
}
