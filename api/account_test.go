package api

import (
	"testing"

	db "github.com/godoquin/simple_bank/db/sqlc"
	"github.com/godoquin/simple_bank/util"
)

func TestGetAccountAPI(t *testing.T) {

}

func ramdomAccount() db.Account {
	return db.Account{
		ID:       util.RamdomInt(1, 1000),
		Owner:    util.RamdomOwner(),
		Balance:  util.RamdomMoney(),
		Currency: util.RamdomCurrency(),
	}
}
