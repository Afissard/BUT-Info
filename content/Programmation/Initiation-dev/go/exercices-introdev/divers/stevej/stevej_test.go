package stevej

import (
	"testing"
)

func TestErr1(t *testing.T) {
	if _, err := ecart(24, 2, 1955); err != nil {
		t.Fail()
	}
}

func TestErr2(t *testing.T) {
	if _, err := ecart(24, 2, 1823); err != errPasUneDate {
		t.Fail()
	}
}

func TestErr3(t *testing.T) {
	if _, err := ecart(24, 2, 3012); err != errPasUneDate {
		t.Fail()
	}
}

func TestErr4(t *testing.T) {
	if _, err := ecart(24, 19, 1955); err != errPasUneDate {
		t.Fail()
	}
}

func TestErr5(t *testing.T) {
	if _, err := ecart(24, 0, 1955); err != errPasUneDate {
		t.Fail()
	}
}

func TestErr6(t *testing.T) {
	if _, err := ecart(32, 3, 1955); err != errPasUneDate {
		t.Fail()
	}
}

func TestErr7(t *testing.T) {
	if _, err := ecart(29, 2, 1955); err != errPasUneDate {
		t.Fail()
	}
}

func TestErr8(t *testing.T) {
	if _, err := ecart(30, 2, 1960); err != errPasUneDate {
		t.Fail()
	}
}

func TestErr9(t *testing.T) {
	if _, err := ecart(31, 4, 1955); err != errPasUneDate {
		t.Fail()
	}
}

func TestErr10(t *testing.T) {
	if _, err := ecart(0, 2, 1955); err != errPasUneDate {
		t.Fail()
	}
}

func TestStevej(t *testing.T) {
	if e, _ := ecart(24, 2, 1955); e != 0 {
		t.Fail()
	}
}

func TestDate1(t *testing.T) {
	if e, _ := ecart(28, 2, 1955); e != 4 {
		t.Fail()
	}
}

func TestDate2(t *testing.T) {
	if e, _ := ecart(12, 2, 1955); e != 12 {
		t.Fail()
	}
}

func TestDate3(t *testing.T) {
	if e, _ := ecart(24, 5, 1955); e != 89 {
		t.Fail()
	}
}

func TestDate4(t *testing.T) {
	if e, _ := ecart(24, 1, 1955); e != 31 {
		t.Fail()
	}
}

func TestDate5(t *testing.T) {
	if e, _ := ecart(24, 2, 1959); e != 1461 {
		t.Fail()
	}
}

func TestDate6(t *testing.T) {
	if e, _ := ecart(24, 2, 1953); e != 730 {
		t.Fail()
	}
}

func TestDate7(t *testing.T) {
	if e, _ := ecart(17, 4, 1986); e != 11375 {
		t.Fail()
	}
}
