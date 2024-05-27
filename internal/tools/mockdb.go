package tools

import (
	"fmt"
)

type mockdb struct{}

var mockLoginDetails = map[string]LoginDetails{
	"juan": {
		Username:  "juan",
		AuthToken: "2121",
	},
	"gabi": {
		Username:  "gabi",
		AuthToken: "0505",
	},
	"lorenzo": {
		Username:  "lorenzo",
		AuthToken: "0101",
	},
}

var mockCoinBalanceDetails = map[string]CoinBalanceDetails{
	"juan": {
		Username:    "juan",
		CoinBalance: 0,
	},
	"gabi": {
		Username:    "gabi",
		CoinBalance: 5000,
	},
	"lorenzo": {
		Username:    "lorenzo",
		CoinBalance: 90000,
	},
}

func (d *mockdb) GetUserLoginDetails(username string) *LoginDetails {
	var loginDetails = LoginDetails{}
	loginDetails, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &loginDetails
}

func (d *mockdb) GetCoinBalanceDetails(username string) *CoinBalanceDetails {
	var coinBalanceDetails = CoinBalanceDetails{}
	coinBalanceDetails, ok := mockCoinBalanceDetails[username]
	if !ok {
		return nil
	}
	return &coinBalanceDetails
}

func (d *mockdb) EditCoinBalanceDetails(username string, operation string, amount int64) *CoinBalanceDetails {
	var coinBalanceDetails = CoinBalanceDetails{}
	coinBalanceDetails, ok := mockCoinBalanceDetails[username]
	if !ok {
		return nil
	}
	fmt.Printf("Received Params and have coinBalance %v, %v, %v, %v \n", username, operation, amount, coinBalanceDetails)

	var newBalance = coinBalanceDetails.CoinBalance
	fmt.Printf("Got newBalance %v \n", newBalance)
	if operation == "+" {
		fmt.Printf("adding \n")
		newBalance += amount
	} else if operation == "-" {
		fmt.Printf("subtracting \n")
		newBalance -= amount
	} else {
		fmt.Printf("error \n")
		return nil
	}
	fmt.Printf("Balance adjusted %v at %v \n", newBalance, &coinBalanceDetails)
	coinBalanceDetails.CoinBalance = newBalance
	mockCoinBalanceDetails[username] = coinBalanceDetails
	fmt.Printf("Mock balance is %v \n", mockCoinBalanceDetails[username])
	return &coinBalanceDetails
}

func (d *mockdb) SetupDatabase() error {
	return nil
}
