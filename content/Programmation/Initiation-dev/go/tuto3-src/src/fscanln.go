package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Ouverture du fichier
	var filePath string = "fichierScan"
	var myFile *os.File
	var err error
	myFile, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Lecture d'une ligne
	var nbLus int
	var unEntier int
	var uneChaine string
	nbLus, err = fmt.Fscanln(myFile, &unEntier, &uneChaine)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("J'ai lu ", nbLus, " valeurs.")
	log.Print("unEntier = ", unEntier)
	log.Print("uneChaine = ", uneChaine)

	// Fermeture du fichier
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
