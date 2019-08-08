package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/pkg/errors"
	"strings"
	"sync"
)

// type to hold any url connection info
type URL struct {
	Scheme string
	Host string
	Port int
	Path string
	Query string
}

// type to hold the app input APIs
type Inputs struct {
	NameURL URL
	JokeURL URL
}

// type to unmarshal name response json
type NameResponse struct {
	First string `json:"name"`
	Last string `json:"surname"`
}

// type to unmarshal joke response json
type JokeResponse struct {
	Status string `json:"type"`
	Value struct {
		Joke string `json:"joke"`
	}
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
func (n *NameResponse) GetNameResponse(u *URL) error {
	body, err := u.RequestInput()
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &n); err != nil {
		msg := fmt.Sprintf("Unable to unmarshal name response JSON from %s. Error: %e", u.Host, err)
		logger.Print(msg)
		return errors.New(msg)
	}

	return nil
}

// make a request to the Joke URL and unmarshal the response body
func (j *JokeResponse) GetJokeResponse(u *URL) error {
	body, err := u.RequestInput()
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &j); err != nil {
		msg := fmt.Sprintf("Unable to unmarshal joke response JSON from %s. Error: %e", u.Host, err)
		logger.Print(msg)
		return errors.New(msg)
	}

	return nil
}

// Invoked by the GetJoke handler, this function coordinates the calls to the
//   dependent name and joke services, then combines the names with the joke and
//   returns the modified joke string.
// TODO: Replace the direct John Doe joke query and replace with a second call
//   from the results of the Name API call. Not sure how to provide the URL.Query
//   field as a parameter that takes inputs. Perhaps using strings.Replace in reverse.
//   Evaluate before doing with Benchmark tests on performance difference, as the
//   WaitGroup would have to be removed.
// TODO: Provide better error handling in the go funcs than panicking.
// TODO: Move the strings.Replace block to its own function for expanded joke
//   formatting, and seek more idiomatic means to replacing the &quot; strings.
func CreateJoke() (string, error) {
	var nr NameResponse
	var jr JokeResponse

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
		if err := jr.GetJokeResponse(&inputs.JokeURL); err != nil {
			logger.Panic(err)
		}
		wg.Done()
	}()
	wg.Wait()

	var joke string
	joke = jr.Value.Joke
	joke = strings.Replace(joke, "John", nr.First, -1)
	joke = strings.Replace(joke, "Doe", nr.Last, -1)
	joke = strings.Replace(joke, "&quot;", "\"", -1)

	return joke, nil
}