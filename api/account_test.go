package api

import db "github.com/b0nbon1/bank-system/db/sqlc"

func TestGetAccountApi(t *testing.T) {
	// Test the GetAccountApi function
	// You can use a mock server or a real server for testing
}

func randomAccount() db.Account {
	return db.Account{
		Username:  randomString(10),
		Password:  randomString(10),
		Email:     randomEmail(),
		FirstName: randomString(5),
		LastName:  randomString(5),
	}
}

