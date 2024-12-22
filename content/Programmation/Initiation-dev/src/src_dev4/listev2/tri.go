package liste

// Trier sorts a list l from the smallest integer to the largest one.
// The sorting is done in place: it will change the list.
func (l *Liste) Trier() {
	ll := parcours(*l, Vide())
	l.contenu = ll.contenu
}

func parcours(l, ll Liste) Liste {
	if l.EstVide() {
		return ll
	}
	lll := insertion(l.Tete(), ll)
	return parcours(l.Queue(), lll)
}

func insertion(v int, l Liste) Liste {
	if l.EstVide() {
		return Concat(v, l)
	}
	if l.Tete() >= v {
		return Concat(v, l)
	}
	ll := insertion(v, l.Queue())
	return Concat(l.Tete(), ll)
}
