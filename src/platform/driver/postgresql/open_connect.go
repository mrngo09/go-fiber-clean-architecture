package driver

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresDB struct {
	SQL *gorm.DB
}

var Postgres = &PostgresDB{}

func ConnectToPostgreSQL() *PostgresDB {
	//host, port, user, password, dbname
	const (
		host     = "localhost"
		port     = "5432"
		user     = "postgres"
		password = ""
		dbname   = "repo-go-gin"
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Logger.LogMode(logger.Info)

	if err != nil {
		log.Fatalln("Cannot connect to Postgresql:", err)
	}
	Postgres.SQL = db
	return Postgres
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
