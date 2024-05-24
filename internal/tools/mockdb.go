package tools

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

func (d *mockdb) SetupDatabase() error {
	return nil
}
