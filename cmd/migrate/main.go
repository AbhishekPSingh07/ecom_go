package main

import (
	"database/sql/driver"
	"log"

	"github.com/AbhishekPSingh07/ecom_go/config"
	"github.com/AbhishekPSingh07/ecom_go/db"
	mySqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func main() {
	db, err := db.NewMySqlStorage(mySqlCfg.Config{
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
	driver,_ := mysql.WithInstance(db,&mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migration",
		"mysql",
		driver,
	)
}
