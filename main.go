package main

import (
	"database/sql"
	"log"

	"github.com/Samarth2898/gobackend/api"
	db "github.com/Samarth2898/gobackend/db/sqlc"
	"github.com/Samarth2898/gobackend/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuation: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
