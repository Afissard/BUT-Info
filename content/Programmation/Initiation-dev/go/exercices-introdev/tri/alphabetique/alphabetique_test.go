package alphabetique

import "testing"

func TestVide1(t *testing.T) {
	dico := make([]string, 0)
	alphabetique(dico)
	if len(dico) != 0 || dico == nil {
		t.Fail()
	}
	alphabetique(nil)
}

func TestVide2(t *testing.T) {
	dico := []string{"bonjour", ""}
	alphabetique(dico)
	if len(dico) != 2 || dico[0] != "" || dico[1] != "bonjour" {
		t.Fail()
	}
	alphabetique(dico)
	if len(dico) != 2 || dico[0] != "" || dico[1] != "bonjour" {
		t.Fail()
	}
}

func TestTri(t *testing.T) {
	dico := []string{
		"boulot", "a", "chat",
		"chateau", "arbre", "bille",
		"armure", "balle", "chignon",
	}
	dicotri := []string{
		"a", "arbre", "armure",
		"balle", "bille", "boulot",
		"chat", "chateau", "chignon",
	}
	alphabetique(dico)
	if !egal(dico, dicotri) {
		t.Fail()
	}
	alphabetique(dico)
	if !egal(dico, dicotri) {
		t.Fail()
	}
}

// Fonctions utilitaires pour les tests
func egal(d, dd []string) bool {
	if len(d) != len(dd) {
		return false
	}
	for i := 0; i < len(d); i++ {
		if d[i] != dd[i] {
			return false
		}
	}
	return true
}
