package main

import (
	"encoding/json"
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

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// Encode the data to JSON, returning the error if there was one
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	// At this point, we know that we won't encounter any more errors before writing the
	// response, so it's safe to add any headers that we want to include. We loop
	// through the header map and add each header to the http.ResponseWriter header map.
	// Note that it's OK if the provided header map is nil. Go doesn't throw an error
	// if you try to range over (or generally, read from) a nil map.
	for key, value := range headers {
		w.Header()[key] = value
	}

	// Add the "Content-Type: application/json" header, then write the status code and
	// JSON response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil

}
