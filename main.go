package main

import (
	"database/sql"
	"log"

	"github.com/etharra/simplebank/api"
	db "github.com/etharra/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:1234@localhost:5432/simple_bank?sslmode=disable"
	serverAdderss = "127.0.0.1:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdderss)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
