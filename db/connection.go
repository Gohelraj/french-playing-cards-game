package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Gohelraj/french-playing-cards-game/config"
	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	// Using a file database for initial project (This can easily be changed later with minimal code changes)
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s", config.Conf.Database.DatabaseFileName))
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := "CREATE TABLE IF NOT EXISTS deck (id text NOT NULL PRIMARY KEY, is_shuffled BOOLEAN, cards json);"
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
	return db
}

