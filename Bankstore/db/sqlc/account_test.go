/*
Pattern: <filename>_test.go
*/

package db

import (
	"context"
	"log"
	"os"
	"testing"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"Bankstore/utils"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://app_user:pswd@localhost:5432/bankdb?sslmode=disable"
)

var ctx = context.Background()

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgx.Connect(ctx, dbSource)
	if err != nil {
		log.Fatal("can not connect to db", err)
	}
	defer conn.Close(ctx)

	testQueries = New(conn)

	os.Exit(m.Run())
}

func TestCreateAccount(t *testing.T){
	createRandomAccount(t)
}

func createRandomAccount(t *testing.T) Account {
	ra := utils.RandomAccount()
	arg := CreateAccountParams {
		Owner: ra.Owner,
		// Balance: utils.RandomInt(0, 1000),
		Balance: ra.Balance,
		Currency: Currency(ra.Currency),
	}
	account, err := testQueries.CreateAccount(ctx, arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	
	return account
}