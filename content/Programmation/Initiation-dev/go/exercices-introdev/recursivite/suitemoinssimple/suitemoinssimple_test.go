package suiteMoinsSimple

import "testing"

func Test1(t *testing.T) {
	res := terme(0)
	if res != 1 {
		t.Error("Le terme u0 est 1, vous dites que c'est", res)
	}
}

func Test2(t *testing.T) {
	res := terme(4)
	if res != 9 {
		t.Error("Le terme u4 est 9, vous dites que c'est", res)
	}
}

// Ajouté après le test 1 de 2022-2023

func Test3(t *testing.T) {
	for i := 0; i < 20; i++ {
		if terme(i) != iteratif(i) {
			t.Fail()
		}
	}
}

func iteratif(n int) (un int) {

	un = 1

	for i := 1; i <= n; i++ {
		if un%2 == 0 {
			un = 3*un + 3
		} else {
			un--
		}
	}

	return
}
