package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/godoquin/simple_bank/db/mock"
	db "github.com/godoquin/simple_bank/db/sqlc"
	"github.com/godoquin/simple_bank/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	account := ramdomAccount()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusCreated, recorder.Code)

}

func ramdomAccount() db.Account {
	return db.Account{
		ID:       util.RamdomInt(1, 1000),
		Owner:    util.RamdomOwner(),
		Balance:  util.RamdomMoney(),
		Currency: util.RamdomCurrency(),
	}
}
