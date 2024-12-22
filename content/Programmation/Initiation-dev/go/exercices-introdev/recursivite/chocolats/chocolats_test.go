package chocolats

import (
	"testing"
)

func TestPauvre(t *testing.T) {
	if miam(3, 4, 2) != 0 {
		t.Fail()
	}
}

func TestRiche(t *testing.T) {
	if miam(1000, 1, 1001) != 1000 {
		t.Fail()
	}
}

func TestTout1(t *testing.T) {
	if miam(16, 2, 2) != 15 {
		t.Fail()
	}
}

func TestTout2(t *testing.T) {
	if miam(15, 1, 3) != 22 {
		t.Fail()
	}
}

func TestTout3(t *testing.T) {
	if miam(1000, 5, 15) != 214 {
		t.Fail()
	}
}

func TestTout4(t *testing.T) {
	if miam(15, 4, 2) != 5 {
		t.Fail()
	}
}
