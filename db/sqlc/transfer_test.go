package db

import (
	"context"
	"testing"
	"time"

	"github.com/eXpect/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	var from_account = createRandomAccount(t)
	var to_account = createRandomAccount(t)

	var arg = CreateTransferParams{
		FromAccountID: from_account.ID,
		ToAccountID:   to_account.ID,
		Amount:        util.RandomMoney(),
	}

	var transfer, err = testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestUpdateTransfer(t *testing.T) {
	var transfer = createRandomTransfer(t)

	var arg = UpdateTransferParams{
		ID:     transfer.ID,
		Amount: util.RandomMoney(),
	}

	var updatedTransfer, err = testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedTransfer)

	require.Equal(t, transfer.ID, updatedTransfer.ID)
	require.Equal(t, transfer.CreatedAt, updatedTransfer.CreatedAt)
	require.Equal(t, transfer.FromAccountID, updatedTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, updatedTransfer.ToAccountID)
	require.Equal(t, arg.Amount, updatedTransfer.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	var transfer = createRandomTransfer(t)

	var deletedTransfer, err = testQueries.DeleteTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, deletedTransfer)

	require.Equal(t, transfer.Amount, deletedTransfer.Amount)
	require.Equal(t, transfer.CreatedAt, deletedTransfer.CreatedAt)
	require.Equal(t, transfer.FromAccountID, deletedTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, deletedTransfer.ToAccountID)
	require.Equal(t, transfer.ID, deletedTransfer.ID)
	require.WithinDuration(t, transfer.CreatedAt, deletedTransfer.CreatedAt, time.Second)
}

func TestGetTransfer(t *testing.T) {
	var transfer = createRandomTransfer(t)

	var existingTransfer, err = testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, existingTransfer)

	require.Equal(t, transfer.ID, existingTransfer.ID)
	require.Equal(t, transfer.FromAccountID, existingTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, existingTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, existingTransfer.Amount)
	require.Equal(t, transfer.CreatedAt, existingTransfer.CreatedAt)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	var arg = ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	var transfers, err = testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
