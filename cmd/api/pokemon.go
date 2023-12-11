package main

import (
	"fmt"
	"net/http"
)

// Add a createPokemonHandler for the "POST /v1/pokemon" endpoint. For now, we simply
// return a plain-text placeholder response.
func (app *application) createPokemonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new pokemon entry")
}

// Add a showMovieHandler for the "GET /v1/pokemon/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.
func (app *application) showPokemonHandler(w http.ResponseWriter, r *http.Request) {

	// grab the ID from the request parameters
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// Otherwise, interpolate the movie ID in a placeholder response.
	fmt.Fprintf(w, "show the details of pokemon %d\n", id)
}
