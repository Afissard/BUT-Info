package main

import (
	//"fmt"
	"log"
	"os"
	//"errors"
	"bufio"
	"strings"
)

type AppData struct {
	err error
	filePath string
	file *os.File
	scanner *bufio.Scanner
}

func (app AppData)Check(){
	if app.err != nil {log.Fatal(app.err)}
}

func (app *AppData)OpenFile(){
	app.file, app.err = os.Open(app.filePath)
	app.Check()
}

func (app *AppData)CloseFile(){
	app.err = app.file.Close()
	app.Check()
}

func (app *AppData) ReadCSV(){
	app.OpenFile()
	app.scanner = bufio.NewScanner(app.file)
	for app.scanner.Scan() {
		log.Print("Je viens de lire : ", app.scanner.Text())
	}
	app.Check()
	app.CloseFile()
}

func(app *AppData) CalcMoy(){
	app.OpenFile()
	for app.scanner.Scan() {
		var line []string = app.scanner.Text()
		var name string = strings.Split(line, ",")
		//var notes int [] = []
		log.Print(name)

		// TODO check bufio and string split
	}

	app.Check()
	app.CloseFile()
}

//////////////////////////////////////////////////
func main(){
	app := AppData{filePath : "notes.csv"}
	//app.ReadCSV()
	app.CalcMoy
	app.Check()
}