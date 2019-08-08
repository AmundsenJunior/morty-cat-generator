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

type NameResponse struct {
	First string `json:"name"`
	Last string `json:"surname"`
}

type JokeResponse struct {
	Status string `json:"type"`
	Value struct {
		Joke string `json:"joke"`
	}
}

// TODO: fix u.Query to not need ? included in URL instantiation in main.go
func (u *URL) SetURL(scheme, host string, port int, path, query string) {
	u.Scheme = scheme
	u.Host = host
	u.Port = port
	u.Path = path
	u.Query = query
}

func (u *URL) GetURL() string {
	return fmt.Sprintf("%s://%s:%d/%s%s", u.Scheme, u.Host, u.Port, u.Path, u.Query)
}

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

//TODO: Replace the direct John Doe joke query and replace with a second call from the results of the Name API call.
// Not sure how to provide the URL.Query field as a parameter that takes inputs. Perhaps using strings.Replace in reverse.
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