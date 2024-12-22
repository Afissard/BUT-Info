package listev1

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
