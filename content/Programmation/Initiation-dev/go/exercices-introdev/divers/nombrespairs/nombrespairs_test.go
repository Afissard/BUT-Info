package nombrespairs

import (
	"testing"
)

func Test1(t *testing.T) {
	if sommeNombresPairs(0, 0) != 0 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if sommeNombresPairs(2, 2) != 2 {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	if sommeNombresPairs(5, 5) != 0 {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	if sommeNombresPairs(0, 10) != 30 {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	if sommeNombresPairs(-10, 0) != -30 {
		t.Fail()
	}
}

func Test6(t *testing.T) {
	if sommeNombresPairs(-10, 10) != 0 {
		t.Fail()
	}
}

func Test7(t *testing.T) {
	if sommeNombresPairs(10, 0) != 30 {
		t.Fail()
	}
}

func Test8(t *testing.T) {
	if sommeNombresPairs(10, -10) != 0 {
		t.Fail()
	}
}
