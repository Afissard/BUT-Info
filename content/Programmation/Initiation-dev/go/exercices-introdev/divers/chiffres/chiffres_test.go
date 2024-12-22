package chiffres

import "testing"

func Test0(t *testing.T) {
	res := chiffres(0)
	if res != 0 {
		t.Error("chiffres(0) devrait retourner 0 mais retourne ", res)
	}
}

func TestPetit(t *testing.T) {
	res := chiffres(3)
	if res != 32 {
		t.Error("chiffres(3) devrait retourner 32 mais retourne ", res)
	}
}

func TestGrand(t *testing.T) {
	res := chiffres(10)
	if res != 405071317 {
		t.Error("chiffres(10) devrait retourner 405071317 mais retourne ", res)
	}
}

// Ajouté suite au test machine 2021-2022
func TestTresGrand(t *testing.T) {
	res := chiffres(25)
	if res != 2284082997 {
		t.Fail()
	}
}
