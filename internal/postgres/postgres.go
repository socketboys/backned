package postgres

import (
	"github.com/jackc/pgx"
	"log"
)

var PostgresConn *pgx.Conn

func InitPostgres() {
	PostgresConn, _ = pgx.Connect(pgx.ConnConfig{
		Host:     "localhost",
		Port:     5433, // will be 5432
		Database: "postgres",
		User:     "postgres",
	})

	if PostgresConn == nil {
		log.Fatal("empty/nil postgres connection established")
	}
}
