package main

import (
	"fmt"
	"net/http"
	"time"

	"pitacca.leoclaudio.dev/internal/data"
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

	user := data.User{
		ID:        userID,
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, user, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "the server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
