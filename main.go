package main

import (
	"database/sql"
	"log"

	"github.com/b0nbon1/bank-system/api"
	db "github.com/b0nbon1/bank-system/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:1234567890@localhost:5432/simple_bank?sslmode=disable"
	serverAddr = ":8080"
)

func main() {
	conn, err :=  sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddr)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	log.Println("server started at", serverAddr)

}