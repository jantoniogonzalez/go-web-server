package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/jantoniogonzalez/go-web-server/api"
	"github.com/jantoniogonzalez/go-web-server/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	// Grab the params and set them on the struct
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	// Instance db interface
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
	// Get User's coin balance
	var tokenDetails *tools.CoinBalanceDetails
	tokenDetails = (*database).GetCoinBalanceDetails(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	// Send Response
	var response = api.SuccessfulResponse{
		Balance: (*tokenDetails).CoinBalance,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

func EditCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var form = api.EditCoinBalanceBody{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	// Get URL params
	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	//Check URL params and Form fields exist
	err = r.ParseForm()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	fmt.Printf("The Form has values %v \n", r.Form)
	// Get Form fields
	err = decoder.Decode(&form, r.Form)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	// Instance Database Interface
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	// Edit Coin Balance Amount
	var coinBalanceDetails *tools.CoinBalanceDetails
	coinBalanceDetails = (*database).EditCoinBalanceDetails(params.Username, form.Operation, form.Amount)
	// Generate Response
	if coinBalanceDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var response = api.SuccessfulResponse{
		Code:    http.StatusOK,
		Balance: coinBalanceDetails.CoinBalance,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
