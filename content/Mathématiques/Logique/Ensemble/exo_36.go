package main

import (
	"fmt"
	"math/rand"
)

type Alien struct {
	Nom     string
	Sexe    string
	Planete string
	Cellule int
}

var Aliens []Alien
var Roswell []Alien
var Roswell_dict = make(map[int]Alien) // Only have Alien in jail

func Create_alien(number int, n string, s string, p string, jail bool) {
	/* Create an Alien and add it to the global dict : Aliens and Roswell */
	var c = 0
	cells++
	if jail {
		c = cells
		Roswell = append(Roswell, Alien{Nom: n, Sexe: s, Planete: p, Cellule: c})
		Roswell_dict[cells] = Roswell[len(Roswell)-1]
	}
	Aliens = append(Aliens, Alien{Nom: n, Sexe: s, Planete: p, Cellule: c})
	//fmt.Println(Aliens[len(Aliens)-1])
}

var cells int = 0
var NOMS = []string{"Albert", "Bob", "Caesar", "Cody", "Bertha", "Elia", "Mathilda", "Aurore"}
var PLANETE = []string{"TERRE", "MARS", "JUPITER", "MERCURE", "NEPTUNE", "PLUTON", "SATURNE", "LUNE", "ANDROMEDE", "PROXIMA_DU_CENTAURE"}
var SEXES = []string{"MALE", "FEMELLE", "AUTRES"}

func Init_random(nbr int) {
	/*Initialise nbr aliens*/
	for i := 0; i < nbr; i++ {
		is_in_jail := false
		rand_temp := rand.Intn(2) // 0 ou 1
		if rand_temp == 1 {
			is_in_jail = false
		} else {
			is_in_jail = true
		}
		Create_alien(nbr, NOMS[rand.Intn(len(NOMS))], SEXES[rand.Intn(len(SEXES))], PLANETE[rand.Intn(len(PLANETE))], is_in_jail)
		fmt.Println(Aliens[i])
	}
}

//-------------------------------------------------------------------------------------//

func ensemble_planete() {
	/*Question 1*/
	var ensemble_P = []string{}
	for x := 0; x < len(Roswell); x++ {
		add_pla := true
		for i := 0; i < len(ensemble_P); i++ {
			if (ensemble_P[i] != Roswell[x].Planete) && add_pla {
				add_pla = true
			} else {
				add_pla = false
			}
		}
		if add_pla {
			ensemble_P = append(ensemble_P, Roswell[x].Planete)
		}
	}
	fmt.Println("Ensemble des planètes des prisoniers : ", ensemble_P)
}

func ensemble_cellules_vide() {
	/*Question 2*/
	var ensemble_nonC = []int{}
	for x := 0; x < cells; x++ {
		if _, is_taken := Roswell_dict[x]; !is_taken {
			ensemble_nonC = append(ensemble_nonC, x)
		}
	}
	fmt.Println("Nombre de cellules vides", len(ensemble_nonC), "\nEnssemble des cellues vides :", ensemble_nonC)
}

func doublons_de_noms() {
	/*Question 3*/
	var ensemble_doublons = []string{}
	for x := 0; x < len(Roswell); x++ {
		for y := 0; y < len(Roswell); y++ {

			if Roswell[x].Nom == Roswell[y].Nom {
				add_double := true
				for i := 0; i < len(ensemble_doublons); i++ {
					if ensemble_doublons[i] == Roswell[x].Nom {
						add_double = false
					}
				}
				if add_double {
					ensemble_doublons = append(ensemble_doublons, Roswell[x].Nom)
				}
			}
		}
	}
	fmt.Println("Ensemble des doublons de noms : ", ensemble_doublons)
}

func main() {
	Init_random(100)
	ensemble_planete()
	ensemble_cellules_vide()
	doublons_de_noms()
}
