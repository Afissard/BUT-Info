package desert

import "testing"

func Test0(t *testing.T) {
	v := duree(0, 0)
	if v != 2 {
		t.Error("Sans eau et sans jamais en trouver, la personne peut survivre 2 jours mais vous indiquez qu'elle peut survivre", v, "jours")
	}
}

func Test1(t *testing.T) {
	v := duree(100, 0)
	if v != 3 {
		t.Error("Avec 100 unités d'eau et sans jamais en trouver, la personne peut survivre 3 jours mais vous indiquez qu'elle peut survivre", v, "jours")
	}
}

func Test2(t *testing.T) {
	v := duree(250, 20)
	if v != 6 {
		t.Error("Avec 250 unités d'eau et en en trouvant 20 par nuit, la personne peut survivre 6 jours mais vous indiquez qu'elle peut survivre", v, "jours")
	}
}

func Test3(t *testing.T) {
	v := duree(500, 30)
	if v != 13 {
		t.Error("Avec 500 unités d'eau et en en trouvant 30 par nuit, la personne peut survivre 13 jours mais vous indiquez qu'elle peut survivre", v, "jours")
	}
}
