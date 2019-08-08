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
			Scheme: "http",
			Host: "localhost",
			Port: 80,
			Path: "uinames/api/",
			Query: "",
		},
		JokeURL: URL{
			Scheme: "http",
			Host: "api.icndb.com",
			Port: 80,
			Path: "jokes/random",
			Query: "firstName=John&lastName=Doe&limitTo=[nerdy]",
		},
	}
}

func main() {
	a := App{}

	a.Initialize()

	a.Run(":8080")
}