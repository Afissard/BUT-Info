package ensemble

import (
	"testing"
)

func TestVide1(t *testing.T) {
	if estEnsemble(nil) {
		t.Fail()
	}
}

func TestVide2(t *testing.T) {
	if !estEnsemble([]int{}) {
		t.Fail()
	}
}

func TestOk(t *testing.T) {
	if !estEnsemble([]int{1, 2, 3, 5, 37, 23, 4, 9}) {
		t.Fail()
	}
}

func TestPasOk(t *testing.T) {
	if estEnsemble([]int{1, 2, 3, 5, 37, 2, 4, 9}) {
		t.Fail()
	}
}
