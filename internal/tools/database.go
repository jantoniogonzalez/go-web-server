package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Username  string
	AuthToken string
}

type CoinBalanceDetails struct {
	CoinBalance int64
	Username    string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetCoinBalanceDetails(username string) *CoinBalanceDetails
	EditCoinBalanceDetails(username string, operation string, amount int64) *CoinBalanceDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {

	var database DatabaseInterface = &mockdb{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
