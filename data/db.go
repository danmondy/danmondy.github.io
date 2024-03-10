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

func RunSql() {
	sql := "ALTER table card add column Image TEXT;"
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func createTables() error {
	//add the schema for my db in code to keep versions in source control
	createCardTableSQL := `CREATE TABLE card (
		"id" TEXT NOT NULL,
		"deckid" TEXT NOT NULL,
		"text" TEXT NOT NULL,
		"image" TEXT NOT NULL,
		"texttopright" INT NOT NULL,
		"textbottomleft" TEXT NOT NULL
	  );`
	createDeckTableSQL := `CREATE TABLE deck (
		"id" TEXT,
		"name" TEXT	
	  );`

	_, err := db.Exec(createCardTableSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = db.Exec(createDeckTableSQL)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func populateTables() {
	decks := []Deck{
		{ID: NewUniqueID(), Name: "Event"},
		{ID: NewUniqueID(), Name: "Item"},
	}
	for _, d := range decks {
		Insert(&d)
	}
	cards := []Card{
		{ID: NewUniqueID(), DeckID: decks[1].ID, Text: "Solar Sails", TextTopRight: 1, TextBottomLeft: "+1 Movement Per Turn"},
		{ID: NewUniqueID(), DeckID: decks[1].ID, Text: "Solar Sails", TextTopRight: 1, TextBottomLeft: "+1 Movement Per Turn"},
		{ID: NewUniqueID(), DeckID: decks[1].ID, Text: "Solar Sails", TextTopRight: 1, TextBottomLeft: "+1 Movement Per Turn"},
		{ID: NewUniqueID(), DeckID: decks[1].ID, Text: "Solar Sails", TextTopRight: 2, TextBottomLeft: "+2 Movement Per Turn"},
		{ID: NewUniqueID(), DeckID: decks[1].ID, Text: "Shield", TextTopRight: 1, TextBottomLeft: "+1 Dmg Absorbed Per Attack"},
		{ID: NewUniqueID(), DeckID: decks[1].ID, Text: "Shield", TextTopRight: 1, TextBottomLeft: "+1 Dmg Absorbed Per Attack"},
		{ID: NewUniqueID(), DeckID: decks[1].ID, Text: "Shield", TextTopRight: 1, TextBottomLeft: "+1 Dmg Absorbed Per Attack"},
		{ID: NewUniqueID(), DeckID: decks[1].ID, Text: "Shield", TextTopRight: 2, TextBottomLeft: "+2 Dmg Absorbed Per Attack"},
		{ID: NewUniqueID(), DeckID: decks[1].ID, Text: "Gatling Gun", TextTopRight: 1, TextBottomLeft: "+1 Movement Per Turn"},
	}
	for _, c := range cards {
		Insert(&c)
	}
}
