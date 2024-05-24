package api

// import (
// 	"encoding/json"
// 	"net/http"
// )

// Here we define the structs of the responses and parameters

// Params
type CoinBalanceParams struct {
	Username string
}

type SuccessfulResponse struct {
	Code int

	Balance int64
}

type Error struct {
	Code int

	Message string
}
