package liste

import "fmt"

// Liste structure defines a standard list. No fields are exported.
type Liste struct {
	contenu *element
}

type element struct {
	v    int
	next *element
}

// Tete allows to access the first element of a list (the head but in french).
// It never modifies the list.
func (l Liste) Tete() int {
	return l.contenu.v
}

// Queue allows to access the liste of all elements but the first of a list (the
// tail but in french).
// It never modifies the list.
func (l Liste) Queue() Liste {
	return Liste{contenu: l.contenu.next}
}

// Concat allows to build a new list from a list and an integer (append but in
// french).
// The integer v is the head of the new list and the list l is its tail.
// It never modifies the list l.
func Concat(v int, l Liste) Liste {
	e := element{v: v, next: l.contenu}
	return Liste{contenu: &e}
}

// Vide returns a new empty list.
func Vide() Liste {
	return Liste{contenu: nil}
}

// EstVide tells whether or not a list l is empty.
func (l Liste) EstVide() bool {
	return l.contenu == nil
}

// Affiche allows to display a list l (as a string).
func (l Liste) Affiche() string {
	var s string
	for !l.EstVide() {
		s += fmt.Sprint(l.Tete(), " ")
		l = l.Queue()
	}
	return s
}
