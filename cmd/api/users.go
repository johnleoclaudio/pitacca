package main

import (
	"fmt"
	"net/http"
)

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `create a new user`)
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	userID, err := app.readUserIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "user details: %s\n", userID)
}
