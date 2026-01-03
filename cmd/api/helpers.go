package main

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readUserIDParam(r *http.Request) (string, error) {
	// When using httprouter, URL parameters are stored in the request context.
	// Use ParamsFromContext() to retrieve a slice containing these parameter names and values
	params := httprouter.ParamsFromContext(r.Context())

	// Use the ByName() method to get the value of the "id" parameter from the slice.
	// /v1/users/{id}
	userID := params.ByName("id")
	if userID == "" {
		return "", errors.New("invalid ID parameter")
	}

	return userID, nil
}
