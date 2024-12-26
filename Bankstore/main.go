package main

import (
	"Bankstore/api"
	db "Bankstore/db/sqlc"
	"context"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgresql://app_user:pswd@localhost:5432/bankdb?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)


func main() {
	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("can not connect to db", err)
	}

	defer pool.Close()

	store := db.NewStore(pool)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Can not start server", err)
	}

}