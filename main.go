package main

import (
	"database/sql"
	"log"

	"github.com/Samarth2898/gobackend/api"
	db "github.com/Samarth2898/gobackend/db/sqlc"
	_ "github.com/lib/pq"
)


const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err!=nil{
		log.Fatal("cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server: ",err)
	}
}