package utils

import (
	"database/sql"
	"log"
	"os"

	"github.com/datahattrick/plusone_someone/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

type Sqldb struct {
	DB *database.Queries
}

var Database Sqldb

func ConnectDB() {
	conn, err := sql.Open("sqlite3", "./sql/db/main.db")
	if err != nil {
		log.Fatal("Failed to connect to database! \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to database successfully")

	db := database.New(conn)

	Database = Sqldb{DB: db}
}
