package data

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func RunSql() {
	/*s := "drop table GameDeck;"
	_, err := db.Exec(s)
	if err != nil {
		fmt.Println(err)
	}
	createGameDeckTable()*/
}

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

// logs fatal if error
func createTables() {
	//add the schema for my db in code to keep versions in source control
	createGameTable()
	createDeckTable()
	createCardTable()
	createBoardTable()
	createGameDeckTable()
	createHexTable()
}

func createCardTable() {
	createTableSQL := `CREATE TABLE card (
		"id" TEXT PRIMARY KEY,
		"deckid" TEXT NOT NULL,
		"text" TEXT NOT NULL,
		"image" TEXT NOT NULL,
		"texttopright" INT NOT NULL,
		"textbottomleft" TEXT NOT NULL,
		"wood" TEXT DEFAULT '0' NOT NULL,
		"ore" TEXT DEFAULT '0' NOT NULL,
		"adder" TEXT DEFAULT '0' NOT NULL,
		"fiber" TEXT DEFAULT '0' NOT NULL
	  );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func createDeckTable() {
	createTableSQL := `CREATE TABLE deck (
		"id" TEXT PRIMARY KEY,
		"name" TEXT	
	  );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
func createGameTable() {
	createTableSQL := `CREATE TABLE game (
		"id" TEXT  PRIMARY KEY,
		"boardid" TEXT,
		"name" TEXT	
	  );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
func createBoardTable() {
	createTableSQL := `CREATE TABLE board (
		"id" TEXT PRIMARY KEY,
		"name" TEXT,
		"colors" TEXT
	  );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func createHexTable() {
	createTableSQL := `CREATE TABLE hex (
		"id" TEXT PRIMARY KEY,		
		"boardid" TEXT NOT NULL,
		"color" TEXT,
		"type" TEXT,
		"x" INT DEFAULT 0 NOT NULL,
		"y" INT DEFAULT 0 NOT NULL
	  );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func createGameDeckTable() {
	createTableSQL := `CREATE TABLE gamedeck (
		"id" TEXT PRIMARY KEY,
		"gameid" TEXT NOT NULL,		
		"deckid" TEXT NOT NULL
	  );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
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
