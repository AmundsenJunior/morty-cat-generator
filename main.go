package main

import (
	"log"
	"os"
)

var logger *log.Logger
var inputs Inputs

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags)

	//TODO: make these values configurable via YAML file
	inputs = Inputs{
		NameURL: URL{
			Scheme: "https",
			Host:   "uinames.com",
			Port:   443,
			Path:   "api/",
			Query:  "",
		},
		FactURL: URL{
			Scheme: "https",
			Host:   "cat-fact.herokuapp.com",
			Port:   443,
			Path:   "facts/random",
			Query:  "",
		},
	}
}

func main() {
	a := App{}

	a.Initialize()

	a.Run(":8080")
}
