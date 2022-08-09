package driver

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	SQL *gorm.DB
}

// postgresql://postgres:root@localhost:5432/
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

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}
	Postgres.SQL = db
	return Postgres
}
