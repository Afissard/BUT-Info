package suiteSimple

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

// Ajouté après le test machine 1 de 2022-2023

func TestGeneral(t *testing.T) {
	for i := 0; i < 100; i++ {
		if terme(i) != calculIteratif(i) {
			t.Fail()
		}
	}
}

func calculIteratif(n int) (un int) {
	un = 1
	for i := 0; i < n; i++ {
		un = (17 * un) % 13
	}
	return
}
