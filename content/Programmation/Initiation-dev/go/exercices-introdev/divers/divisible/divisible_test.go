package divisible

import "testing"

func Test1(t *testing.T) {
	if estDivisible(5, 4) {
		t.Error("Vous affirmez que 5 est divisible par 4 ou que 4 est divisible par 5, ce qui n'est pas vrai")
	}
}

func Test2(t *testing.T) {
	if !estDivisible(2, 4) {
		t.Error("Vous affirmez que 4 n'est pas divisible par 2, ce qui n'est pas vrai")
	}
}

// Ajouté après le test 1 de 2022-2023

func Test3(t *testing.T) {
	if estDivisible(11, 98) {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	if !estDivisible(79560, 312) {
		t.Fail()
	}
}
