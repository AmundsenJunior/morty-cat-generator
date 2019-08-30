package main

import (
	"encoding/json"
	"fmt"
	rnm "github.com/pitakill/rickandmortyapigowrapper"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

type URL struct {
	Scheme string
	Host   string
	Port   int
	Path   string
	Query  string
}

type CatFactsAPI struct {
	Url URL
}

type RickAndMortyAPI struct  {
	Url URL
	Info struct {
		CharactersNumber int `json:"count"`
	} `json:"info"`
}

type CatFact struct {
	Fact string `json:"text"`
}

type MortyName struct {
	Character *rnm.Character
}

// return the formatted string of a URL
func (u *URL) getURL() string {
	return fmt.Sprintf("%s://%s:%d/%s?%s", u.Scheme, u.Host, u.Port, u.Path, u.Query)
}

// make a generic GET request for the URL and return raw []byte response body
//TODO: IDE reports an unhandled error on the defer statement
func (c *CatFactsAPI) GetCatFactResponse() ([]byte, error) {
	response, err := http.Get(c.Url.getURL())
	if err != nil {
		msg := fmt.Sprintf("Failed to make request to %s. Error: %e", c.Url.Host, err)
		return nil, errors.New(msg)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		msg := fmt.Sprintf("Unable to read response body from %s. Error: %e", c.Url.Host, err)
		return nil, errors.New(msg)
	}

	return body, nil
}

// make a request to the Fact URL and unmarshal the response body
func (f *CatFact) GetCatFact(c *CatFactsAPI) error {
	body, err := c.GetCatFactResponse()
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &f); err != nil {
		msg := fmt.Sprintf("Unable to unmarshal name response JSON from %s. Error: %e", c.Url.Host, err)
		logger.Print(msg)
		return errors.New(msg)
	}

	return nil
}

// get number of characters available in the Rick and Morty API /character endpoint
// takes n default value if requesting the api errors
//TODO: IDE reports an unhandled error on the defer statement
func (r *RickAndMortyAPI) GetNumberOfMortyNames(n int) {
	response, err := http.Get(r.Url.getURL())
	if err != nil {
		logger.Printf("Failed to make request to %s. Error: %e", r.Url.Host, err)
		r.Info.CharactersNumber = n
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Printf("Unable to read response body from %s. Error: %e", r.Url.Host, err)
		r.Info.CharactersNumber = n
	}

	if err = json.Unmarshal(body, &r); err != nil {
		logger.Printf("Unable to unmarshal name response JSON from %s. Error: %e", r.Url.Host, err)
		r.Info.CharactersNumber = n
	}
}

// make a request to the Name URL and unmarshal the response body
// TODO: Check for empty name in rnm api response and retry
func (mn *MortyName) GetMortyName(r *RickAndMortyAPI) {
	rando := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomcharacterid := rando.Intn(r.Info.CharactersNumber)

	mn.Character, _ = rnm.GetCharacter(randomcharacterid)
}

// Invoked by the GetFact handler, this function coordinates the calls to the
//   dependent name and fact services, then combines the names with the fact and
//   returns the modified fact string.
// TODO: Provide better error handling in the go funcs than panicking.
// TODO: Enable gender queries to rnm api based on pronouns used in cat fact.
// TODO: If no name replacements are made, retrieve a new fact.
func CreateMortyFact() (string, error) {
	var mn MortyName
	var cf CatFact

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		mn.GetMortyName(&rickandmortyapi)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		if err := cf.GetCatFact(&catfactsapi); err != nil {
			logger.Panic(err)
		}
		wg.Done()
	}()
	wg.Wait()

	// TODO: Move the strings.Replace block to its own function for expanded fact formatting.
	var fact string
	fact = cf.Fact

	fact = strings.Replace(fact, "cat ", fmt.Sprintf("%s ", mn.Character.Name), -1)
	fact = strings.Replace(fact, "Cat ", fmt.Sprintf("%s ", mn.Character.Name), -1)
	fact = strings.Replace(fact, "cat.", fmt.Sprintf("%s.", mn.Character.Name), -1)
	fact = strings.Replace(fact, "cat's", fmt.Sprintf("%s's", mn.Character.Name), -1)
	fact = strings.Replace(fact, "cats", fmt.Sprintf("%ss", mn.Character.Name), -1)
	fact = strings.Replace(fact, "Cats", fmt.Sprintf("%ss", mn.Character.Name), -1)
	fact = strings.Replace(fact, "Kittens", fmt.Sprintf("%ss", mn.Character.Name), -1)
	fact = strings.Replace(fact, "kittens", fmt.Sprintf("%ss", mn.Character.Name), -1)
	fact = strings.Replace(fact, "kitten", mn.Character.Name, -1)

	return fact, nil
}
