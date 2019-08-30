package main

import "net/http"

// TODO: unhandled error on w.Write call
func (a *App) GetMortyFact(w http.ResponseWriter, r *http.Request) {
	fact, err := CreateMortyFact()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	w.Write([]byte(fact))
}

// TODO: check server availability of Name and Joke URLs
// TODO: return JSON response of 200 with {status:[name:<red or green, fact:<red or green>]}
func (a *App) GetHealthStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
