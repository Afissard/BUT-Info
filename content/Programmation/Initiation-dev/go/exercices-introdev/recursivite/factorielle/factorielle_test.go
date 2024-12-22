package factorielle

import (
	"testing"
)

func Test1(t *testing.T) {
	if factorielle(0) != 1 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if factorielle(5) != 120 {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	if factorielle(10) != 3628800 {
		t.Fail()
	}
}
