package main

import (
	"fmt"
	"log"
	"os"
	//"io/ioutil"
	//"errors"
)

type App_data struct {
	err error
	file *os.File
	file_name string
}

func main() {
	app := App_data{file_name : "fichier"}

	app.file, app.err = os.Create(app.file_name)
	if app.err != nil {
		log.Fatal(app.err)
	}

	_, app.err = fmt.Fprintln(app.file, "Bonjour")
	if app.err != nil {
		log.Fatal(app.err)
	}

	app.Mult()

	_, app.err = fmt.Fprintln(app.file, "Au revoir")
	if app.err != nil {
		log.Fatal(app.err)
	}

	app.Fread()

	app.err = app.file.Close()
	if app.err != nil {
		log.Fatal(app.err)
	}
}
