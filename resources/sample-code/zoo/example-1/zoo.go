package main

import (
	"os"
	"path"
	"runtime"
	"text/template"
	"time"
)

type Conf struct {
	TimeGenerated string // Time a Configuration was Created
	Zoos          []Zoo  // List of Zoos
}

type Zoo struct {
	Name    string   // Name of Zoo
	Climate string   // Climate of Area where Zoo is located
	Animals []Animal // Animals in Zoo
}

type Animal struct {
	Name     string   // Name of the Animal
	Climates []string // Climates where it can live
}

var (
	// Animals
	Alligator = Animal{Name: "Alligator", Climates: []string{"Tropical", "SubTropical"}}
	Crocodile = Animal{Name: "Crocodile", Climates: []string{"Tropical", "SubTropical"}}
	ArcticFox = Animal{Name: "ArcticFox", Climates: []string{"Arctic", "SubArctic"}}
	Puffin    = Animal{Name: "Puffin", Climates: []string{"Arctic", "SubArctic"}}

	// Zoos
	miamiZoo    = Zoo{Name: "MiamiZoo", Climate: "SubTropical", Animals: []Animal{Alligator, Puffin}}
	reykjavikZoo = Zoo{Name: "ReykjavikZoo", Climate: "SubArctic", Animals: []Animal{ArcticFox, Crocodile}}
)

func main() {
	conf := Conf{TimeGenerated: time.Now().UTC().String(), Zoos: []Zoo{miamiZoo, reykjavikZoo}}

	// Gets the directory this file is in
	// See: https://stackoverflow.com/questions/32163425/golang-how-to-get-the-directory-of-the-package-the-file-is-in-not-the-current-w
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	// The template name must match a file in ParseFiles which is the main template
	// See: https://stackoverflow.com/questions/10199219/go-template-function
	tmpl, err := template.New("zoo.tmpl").ParseFiles(path.Dir(filename) + "/zoo.tmpl")
	if err != nil {
		panic(err)
	}

	// Here we use the template and conf to make to generate textual output
	// We are using 'os.Stdout` to output to screen, a file can be used instead
	err = tmpl.Execute(os.Stdout, conf)
	if err != nil {
		panic(err)
	}

}
