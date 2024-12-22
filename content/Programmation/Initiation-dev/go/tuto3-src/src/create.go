package main

import (
	"log"
	"os"
)

func main() {
	var myFile *os.File
	var err error
	myFile, err = os.Create("monfichier")
	if err != nil {
		log.Fatal(err)
	}
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
