package double

import "testing"

func TestZero(t *testing.T) {
	x := 0
	double(&x)
	if x != 0 {
		t.Error("double(&x) avec x = 0 devrait donner x = 0",
			"mais donne x=", x)
	}
}

func TestDix(t *testing.T) {
	x := 10
	double(&x)
	if x != 20 {
		t.Error("double(&x) avec x = 10 devrait donner x = 20",
			"mais donne x=", x)
	}
}

func TestMoinsCinq(t *testing.T) {
	x := -5
	double(&x)
	if x != -10 {
		t.Error("double(&x) avec x = -5 devrait donner x = -10",
			"mais donne x=", x)
	}
}
