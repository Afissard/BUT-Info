package main

import (
	"fmt"
	"log"
	//"os"
	"io/ioutil"
	"errors"
)

func (app *App_data) Fread() {
	var data []byte
	data, app.err = ioutil.ReadFile(app.file_name)
	if app.err != nil {
		log.Fatal(app.err)
	}
	log.Print("J'ai lu :\n", string(data))
}

func (app *App_data) Mult(){
	fmt.Print("Table de : ")
    var input int
    _, app.err = fmt.Scanf("%d", &input)
	
	if input>10 {
		app.err = errors.New("FATAL ERROR : Input higher than 10.") // custom error
		log.Fatal(app.err)
	}

	for i:=0; i<10; i++{
		var mult_i int = input*i
		_, app.err = fmt.Fprintln(app.file, mult_i)
		if app.err != nil {
			log.Fatal(app.err)
		}
	}
}