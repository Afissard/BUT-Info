package suite

import "testing"

func TestSuite(t *testing.T) {
	x := terme(1000)
	if x != 493 {
		t.Error("U(1000) vaut 493 mais vous indiquez qu'il vaut", x)
	}
}
