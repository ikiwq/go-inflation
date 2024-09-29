package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitSqlClient(driver string, username string, password, host string, dbname string) *sqlx.DB {
	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", username, password, host, dbname)
	sqlClient, err := sqlx.Connect(driver, connStr)
	if err != nil {
		log.Fatalf("error while connecting to SQL database: %v", err)
	}

	if err := sqlClient.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("successfully connected to the SQL database")
	return sqlClient
}
