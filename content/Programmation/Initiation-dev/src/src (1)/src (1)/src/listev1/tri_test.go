package listev1

import (
	"math/rand"
	"testing"
)

func TestTri(t *testing.T) {
	l := Vide()
	l.Trier()
	if !l.EstVide() {
		t.Error(l.Affiche())
	}

	l = listeAleatoire(10)
	l.Trier()
	if !l.estTriee() {
		t.Error(l.Affiche())
	}

	l.Trier()
	if !l.estTriee() {
		t.Error(l.Affiche())
	}
}

func BenchmarkTri(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := listeAleatoire(2500)
		l.Trier()
	}
}

func listeAleatoire(length int) Liste {
	l := Vide()
	for i := 0; i < length; i++ {
		v := rand.Intn(100)
		l = Concat(v, l)
	}
	return l
}

func (l Liste) estTriee() bool {
	dernier := l.Tete()
	l = l.Queue()
	for !l.EstVide() {
		if dernier > l.Tete() {
			return false
		}
		dernier = l.Tete()
		l = l.Queue()
	}
	return true
}
