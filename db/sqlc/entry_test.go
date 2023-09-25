package db

import (
	"context"
	"testing"
	"time"

	"github.com/eXpect/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	var account = createRandomAccount(t)
	var arg = CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	var entry, err = testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestUpdateEntry(t *testing.T) {
	var entry = createRandomEntry(t)

	var arg = UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}

	var updatedEntry, err = testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)

	require.Equal(t, entry.ID, updatedEntry.ID)
	require.Equal(t, entry.AccountID, updatedEntry.AccountID)
	require.Equal(t, arg.Amount, updatedEntry.Amount)
	require.Equal(t, entry.CreatedAt, updatedEntry.CreatedAt)
	require.WithinDuration(t, entry.CreatedAt, updatedEntry.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	var entry = createRandomEntry(t)

	var deletedEntry, err = testQueries.DeleteEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, deletedEntry)

	require.Equal(t, entry.AccountID, deletedEntry.AccountID)
	require.Equal(t, entry.Amount, deletedEntry.Amount)
	require.Equal(t, entry.CreatedAt, deletedEntry.CreatedAt)
	require.Equal(t, entry.ID, deletedEntry.ID)
}

func TestGetEntry(t *testing.T) {
	var entry = createRandomEntry(t)

	var existingEntry, err = testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, existingEntry)

	require.Equal(t, entry.ID, existingEntry.ID)
	require.Equal(t, entry.AccountID, existingEntry.AccountID)
	require.Equal(t, entry.CreatedAt, existingEntry.CreatedAt)
	require.Equal(t, entry.Amount, existingEntry.Amount)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	var arg = ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	var entries, err = testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
