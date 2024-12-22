package nombrepremier

import (
	"testing"
)

func Test1(t *testing.T) {
	if estPremier(0) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if estPremier(1) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	if !estPremier(2) {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	if !estPremier(3) {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	if estPremier(4) {
		t.Fail()
	}
}

func Test6(t *testing.T) {
	if !estPremier(102647) {
		t.Fail()
	}
}

func Test7(t *testing.T) {
	if estPremier(102647032) {
		t.Fail()
	}
}
