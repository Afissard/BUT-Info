package main

import (
	"fmt"
	"math"
)

// Interface
type Forme interface {
	Perimetre() float64
}

func PlusGrandPerimetre(f1, f2 Forme) bool {
	return f1.Perimetre() > f2.Perimetre()
}

// Implantation pour les carrés
type Carre struct {
	Cote float64
}

func (c Carre) Perimetre() float64 {
	return 4 * c.Cote
}

// Implantation pour les cercles
type Cercle struct {
	Rayon float64
}

func (c Cercle) Perimetre() float64 {
	return 2 * math.Pi * c.Rayon
}

func main() {
	var ca Carre = Carre{Cote: 5}
	var ce Cercle = Cercle{Rayon: 2.5}
	fmt.Println(PlusGrandPerimetre(ca, ce))
}
