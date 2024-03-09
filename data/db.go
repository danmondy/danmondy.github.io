package data

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func Initialize(recreate bool) error {
	if recreate {
		DeleteAndRecreateFile()
	}

	var err error
	db, err = sqlx.Open("sqlite3", "./card.db")
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return err
	}

	if recreate {
		createTables()
		populateTables()
	}

	return nil
}

func OpenNew(fileName string) error {
	db.Close()
	var err error
	db, err = sqlx.Open("sqlite3", fmt.Sprintf("./%v", fileName))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func DeleteAndRecreateFile() {
	os.Remove("card.db") // delete the file

	log.Println("Creating sqlite-database...")
	file, err := os.Create("card.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("...created.")
}
