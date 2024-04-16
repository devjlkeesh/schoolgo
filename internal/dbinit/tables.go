package dbinit

import (
	"database/sql"
	"log"
)

func InitUserTable(db *sql.DB) {
	log.Println("1 initializing")
	query := `CREATE TABLE IF NOT EXISTS USERS (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
			"username" TEXT,
			"password" TEXT,
			"role" TEXT,
			"status" TEXT,
			"lastLoginAt" TEXT
		  );`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func InitTeacherTable(db *sql.DB) {
	log.Println("2 initializing")
	query := `CREATE TABLE IF NOT EXISTS TEACHERS (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
			"fullname" TEXT,
			"age" integer,
			"birthdate" TEXT,
			"subject" TEXT,
			"category" TEXT,
			"phone" TEXT,
			"img" TEXT,
			"isImgPublic" TEXT
		  );`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func InitNewsTable(db *sql.DB) {
	log.Println("3 initializing")
	query := `CREATE TABLE IF NOT EXISTS NEWS (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
			"title" TEXT,
			"overview" TEXT,
			"body" TEXT,
			"isNews" TEXT,
			"createdAt" TEXT,
			"createdAtMilli" integer,
			"createdBy" TEXT,
			"accepted" TEXT
		  );`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}
