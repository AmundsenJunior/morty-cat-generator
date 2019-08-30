package main

import (
	"log"
	"os"
)

var a App
var logger *log.Logger
var catfactsapi CatFactsAPI
var rickandmortyapi RickAndMortyAPI

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags)

	catfactsapi = CatFactsAPI{
		Url: URL{
			Scheme: "https",
			Host:   "cat-fact.herokuapp.com",
			Port:   443,
			Path:   "facts/random",
			Query:  "",
		},
	}

	// This configuration is only used to make the custom call to get the count of all characters. Otherwise, all
	// interactions with the Rick and Morty API are done via the API Go wrapper package.
	rickandmortyapi = RickAndMortyAPI{
		Url: URL{
			Scheme: "https",
			Host: "rickandmortyapi.com",
			Port: 443,
			Path: "api/character/",
			Query: "",
		},
	}

	rickandmortyapi.GetNumberOfMortyNames(100)

	a = App{}
	a.Initialize()

}

func main() {
	a.Run(":8000")
}
