package database

import (
	"database/sql"
)

var DB *sql.DB

func SetDb(inDB *sql.DB) {
	DB = inDB
}
