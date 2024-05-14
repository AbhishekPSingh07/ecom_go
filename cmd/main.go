package main

import (
	"database/sql"
	"log"

	"github.com/AbhishekPSingh07/ecom_go/cmd/api"
	"github.com/AbhishekPSingh07/ecom_go/config"
	"github.com/AbhishekPSingh07/ecom_go/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySqlStorage(mysql.Config{
		User:                 config.Envs.DBuser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully Connected!")
}