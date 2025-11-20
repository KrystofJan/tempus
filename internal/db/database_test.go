package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/KrystofJan/tempus/internal/constants"
)

func NewTestDatabase() (*Database, error) {
	dbPath, err := getDBPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Println("Did not find the database, creating one...")
	}

	db, err := sql.Open(constants.DATABASE_DRIVER, dbPath)
	if err != nil {
		return nil, err
	}

	return &Database{
		Instance: db,
	}, nil
}

func GetTestConnString() (string, error) {
	const dbPath = ""
	return fmt.Sprintf("sqlite3://%s", dbPath), nil
}
