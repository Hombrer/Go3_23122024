package db

import (
	"Bankstore/utils"
	"context"
	"testing"
	"time"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams(utils.RandomUser())
	
}