package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// type to hold any url connection info
type URL struct {
	Scheme string
	Host   string
	Port   int
	Path   string
	Query  string
}

// type to hold the app input APIs
type Inputs struct {
	NameURL URL
	FactURL URL
}

// type to unmarshal name response json
type NameResponse struct {
	First string `json:"name"`
	Last  string `json:"surname"`
}

// type to unmarshal joke response json
type FactResponse struct {
	Fact string `json:"text"`
}

// take input of the components of each dependent service's url
//TODO: add input validation here of url configuration values
func (u *URL) SetURL(scheme, host string, port int, path, query string) {
	u.Scheme = scheme
	u.Host = host
	u.Port = port
	u.Path = path
	u.Query = query
}

// return the formatted string of a URL
func (u *URL) GetURL() string {
	return fmt.Sprintf("%s://%s:%d/%s?%s", u.Scheme, u.Host, u.Port, u.Path, u.Query)
}

// make a generic GET request for the URL and return raw []byte response body
//TODO: IDE reports an unhandled error on the defer statement
func (u *URL) RequestInput() ([]byte, error) {
	response, err := http.Get(u.GetURL())
	if err != nil {
		msg := fmt.Sprintf("Failed to make request to %s. Error: %e", u.Host, err)
		return nil, errors.New(msg)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		msg := fmt.Sprintf("Unable to read response body from %s. Error: %e", u.Host, err)
		return nil, errors.New(msg)
	}

	return body, nil
}

// make a request to the Name URL and unmarshal the response body
func GetResponse(u *URL, r interface{}) error {
	body, err := u.RequestInput()
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &r); err != nil {
		msg := fmt.Sprintf("Unable to unmarshal name response JSON from %s. Error: %e", u.Host, err)
		logger.Print(msg)
		return errors.New(msg)
	}

	return nil
}

// make a request to the Name URL and unmarshal the response body
func (n *NameResponse) GetNameResponse(u *URL) error {
	err := GetResponse(u, n)
	if err != nil {
		return err
	}

	return nil
}

// make a request to the Fact URL and unmarshal the response body
func (f *FactResponse) GetFactResponse(u *URL) error {
	err := GetResponse(u, f)
	if err != nil {
		return err
	}

	return nil
}

// Invoked by the GetFact handler, this function coordinates the calls to the
//   dependent name and fact services, then combines the names with the fact and
//   returns the modified fact string.
// TODO: Provide better error handling in the go funcs than panicking.
// TODO: Move the strings.Replace block to its own function for expanded fact formatting.
func CreateFact() (string, error) {
	var nr NameResponse
	var fr FactResponse

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := nr.GetNameResponse(&inputs.NameURL); err != nil {
			logger.Panic(err)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		if err := fr.GetFactResponse(&inputs.FactURL); err != nil {
			logger.Panic(err)
		}
		wg.Done()
	}()
	wg.Wait()

	var fact string
	fact = fr.Fact
	fact = strings.Replace(fact, "cat ", fmt.Sprintf("%s ", nr.First), -1)
	fact = strings.Replace(fact, "cat.", fmt.Sprintf("%s.", nr.First), -1)
	fact = strings.Replace(fact, "cat's", fmt.Sprintf("%s's", nr.First), -1)
	fact = strings.Replace(fact, "cats", fmt.Sprintf("%ss", nr.Last), -1)
	fact = strings.Replace(fact, "Cats", fmt.Sprintf("%s %ss", nr.First, nr.Last), -1)

	return fact, nil
}
