package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"school23/internal/database"
	"school23/internal/dbinit"
	"school23/internal/routes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := routes.NewRouter()

	dbFileName := "./sqlite-database.db"
	createDbFileIfNotExists(dbFileName)

	sqliteDatabase, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		log.Fatalln("can not connect to database", err)
	}
	defer sqliteDatabase.Close()

	dbinit.InitUserTable(sqliteDatabase)
	dbinit.InitTeacherTable(sqliteDatabase)
	dbinit.InitNewsTable(sqliteDatabase)

	database.SetDb(sqliteDatabase)

	err = http.ListenAndServe(":8372", router)
	if err != nil {
		panic(err)
	}
}

func createDbFileIfNotExists(filename string) {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(filename)
			if err != nil {
				log.Fatalln("can not connect database ", err.Error())
			}
			defer f.Close()
		} else {
			log.Fatalln(err)
		}
	}
}
