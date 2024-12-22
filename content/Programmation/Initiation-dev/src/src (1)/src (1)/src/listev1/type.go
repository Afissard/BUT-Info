package listev1

import "fmt"

type Liste struct {
	contenu []int
}

func (l Liste) Tete() int {
	return l.contenu[0]
}

func (l Liste) Queue() Liste {
	contenu := make([]int, len(l.contenu)-1)
	copy(contenu, l.contenu[1:])
	return Liste{contenu: contenu}
}

func Concat(v int, l Liste) Liste {
	contenu := make([]int, len(l.contenu)+1)
	copy(contenu[1:], l.contenu)
	contenu[0] = v
	return Liste{contenu: contenu}
}

func Vide() Liste {
	return Liste{contenu: []int{}}
}

func (l Liste) EstVide() bool {
	return len(l.contenu) == 0
}

func (l Liste) Affiche() string {
	var s string
	for !l.EstVide() {
		s += fmt.Sprint(l.Tete(), " ")
		l = l.Queue()
	}
	return s
}
