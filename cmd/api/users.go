package main

import (
	"fmt"
	"net/http"
	"time"

	"pitacca.leoclaudio.dev/internal/data"
)

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `create a new user`)

	var input struct {
		Email string `json:"email"`
	}
	err := app.readJSON(r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	userID, err := app.readUserIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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
		app.serverErrorResponse(w, r, err)
	}
}
