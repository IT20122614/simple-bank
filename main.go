package main

import (
	"database/sql"
	"log"

	"github.com/IT20122614/simple-bank/api"
	db "github.com/IT20122614/simple-bank/db/sqlc"
	"github.com/IT20122614/simple-bank/util"
	_ "github.com/lib/pq"
)



func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServiceAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
