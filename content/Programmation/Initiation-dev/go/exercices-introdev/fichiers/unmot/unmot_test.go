package unmot

import "testing"

func TestPasLa(t *testing.T) {
	_, err := premiermot("test")
	if err != errImpossible {
		t.Fail()
	}
}

// ajoutés après le test machine

func TestVide(t *testing.T) {
	_, err := premiermot("../fichiers-tests/unmot-test2")
	if err != errImpossible {
		t.Fail()
	}
}

func TestOk(t *testing.T) {
	mot, err := premiermot("../fichiers-tests/unmot-test3")
	if err != nil {
		t.Fail()
		return
	}
	if mot != "bonjour" {
		t.Fail()
	}
}

func TestOk2(t *testing.T) {
	mot, err := premiermot("../fichiers-tests/unmot-test4")
	if err != nil {
		t.Fail()
		return
	}
	if mot != "saucisse" {
		t.Fail()
	}
}
