package db

import (
	"Bankstore/utils"
	"context"
	"testing"
	"time"

	// "time"
	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, acc Account) Entry {
	arg := CreateEntryParams {
		AccountID: acc.ID,
		Amount: utils.RandomInt(1, 100),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	acc := createRandomAccount(t)
	createRandomEntry(t, acc)
}

func TestGetEntry(t *testing.T) {
	acc := createRandomAccount(t)
	entry1 := createRandomEntry(t, acc)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)

	require.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)

}

func TestUpdateEntry(t *testing.T) {
	acc := createRandomAccount(t)
	entry1 := createRandomEntry(t, acc)

	arg := UpdateEntryParams {
		ID: entry1.ID,
		Amount: utils.RandomInt(1, 30),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, arg.Amount, entry2.Amount)

	require.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)

}

func TestDeleteEntry(t *testing.T) {
	acc := createRandomAccount(t)
	entry1 := createRandomEntry(t, acc)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)

	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.Empty(t, entry2)

	require.EqualError(t, err, pgx.ErrNoRows.Error())
}

func TestListEntries(t *testing.T) {
	acc := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, acc)
	}

	arg := ListEntriesParams {
		AccountID: acc.ID,
		Limit: 5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}