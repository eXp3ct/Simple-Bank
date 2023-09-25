package db

import (
	"context"
	"testing"
	"time"

	"github.com/eXpect/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	var arg = CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	var account, err = testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	var account = createRandomAccount(t)

	var existingAccount, err = testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, existingAccount)

	require.Equal(t, account.ID, existingAccount.ID)
	require.Equal(t, account.Balance, existingAccount.Balance)
	require.Equal(t, account.CreatedAt, existingAccount.CreatedAt)
	require.Equal(t, account.Currency, existingAccount.Currency)
	require.Equal(t, account.Owner, existingAccount.Owner)

}

func TestUpdateAccount(t *testing.T) {
	var account = createRandomAccount(t)

	var arg = UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	var updatedAccount, err = testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, account.CreatedAt, updatedAccount.CreatedAt)
	require.Equal(t, account.Currency, updatedAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, updatedAccount.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	var account = createRandomAccount(t)

	var deletedAccount, err = testQueries.DeleteAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, deletedAccount)

	require.Equal(t, account.ID, deletedAccount.ID)
	require.Equal(t, account.Balance, deletedAccount.Balance)
	require.Equal(t, account.Owner, deletedAccount.Owner)
	require.Equal(t, account.CreatedAt, deletedAccount.CreatedAt)
	require.Equal(t, account.Currency, deletedAccount.Currency)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	var arg = ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	var accounts, err = testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
