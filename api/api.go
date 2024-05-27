package api

import (
	"encoding/json"
	"net/http"
)

// Here we define the structs of the responses and parameters

// Params
type CoinBalanceParams struct {
	Username string
}

// Reponse
type SuccessfulResponse struct {
	Code int

	Balance int64
}

type Error struct {
	Code int

	Message string
}

// Body
type EditCoinBalanceBody struct {
	Username string

	Operation string

	Amount int64
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	// Set Headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// Set response
	json.NewEncoder(w).Encode(resp)
}

var (
	// if there is an error with the request
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	// if there is an error with our code, we send this
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error had occured", http.StatusInternalServerError)
	}
)
