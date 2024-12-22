package main

import (
	"fmt"
	"log"
	"os"
)

func (app *App_data) Fget() {
	// Ouverture du fichier
	var app.err app.error
	app.file, app.err = os.Open(app.file_name)
	if app.err != nil {
		log.Fatal(app.err)
	}

	// Lecture d'une ligne
	var nbLus int
	var unEntier int
	var uneChaine string
	nbLus, app.err = fmt.Fscanln(app.file, &unEntier, &uneChaine)
	if app.err != nil {
		log.Fatal(app.err)
	}
	log.Print("J'ai lu ", nbLus, " valeurs.")
	log.Print("unEntier = ", unEntier)
	log.Print("uneChaine = ", uneChaine)

	// Fermeture du fichier
	app.err = app.file.Close()
	if app.app.err != nil {
		log.Fatal(app.err)
	}
}
