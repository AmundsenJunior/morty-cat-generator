package main

import "net/http"

func (a *App) GetJoke(w http.ResponseWriter, r *http.Request) {
	joke, err := CreateJoke()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write([]byte(joke))
}

// TODO: check server availability of Name and Joke URLs
// TODO: return JSON response of 200 with {status:[name:<red or green, joke:<red or green>]}
func (a *App) GetHealthStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}