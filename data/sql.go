package data

import (
	"log"
)

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

func DeleteCardsInDeck(id string) error {
	_, err := db.Exec("delete from card where deckid = ?", id)
	return err
}

/*
func GetCardsForDeck(id string) ([]Card, error) {
	cards := make([]Card, 0)
	err := db.Select(&cards, "SELECT c.* FROM card c join deckcard dc on c.id = dc.cardid where dc.deckid = ?", id)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func AddCardToDeck(deckID string, cardID string) {
	dc := DeckCard{DeckID: deckID, CardID: cardID}
	Insert(&dc)
}*/

//use reflect instead
/*func UpdateEnv(env Envelope) error {
	log.Println("Updating envelope...", env)
	_, err := db.NamedExec("UPDATE envelope set name = :name, balance = :balance where id = :id;", env)
	log.Println("Envlope complete", env)
	return err
}

// ----ENVELOPES----
func GetAllEnvelopes() ([]Envelope, error) {
	envs := []Envelope{}
	fmt.Println("getting all envelopes")
	err := db.Select(&envs, "SELECT * FROM envelope")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(envs)
	return envs, nil
}
func UpdateEnvelope(e *Envelope) error {
	fmt.Println("Updating:", e, "...")
	_, err := db.NamedExec("UPDATE envelope set name = :name, balance = :balance, budget = :budget, frequency = :frequency where id = :id;", e)
	return err
}

// ----ACCOUNTS----
func GetAllAccounts() ([]Account, error) {
	accounts := []Account{}
	fmt.Println("getting all accounts")
	err := db.Select(&accounts, "SELECT * FROM account")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(accounts)
	return accounts, nil
}

func UpdateAccount(a *Account) error {
	fmt.Println("Updating:", a, "...")
	_, err := db.NamedExec("UPDATE account set name = :name, balance = :balance where id = :id;", a)
	return err
}

// ----TRANSACTIONS----
// ID could be for an envelope or an account
func GetAllTransactionsForID(id string) ([]Tx, error) {

	tx := []Tx{}

	err := db.Select(&tx, "SELECT *from tx where fromid = ? OR toid = ? order by date", id, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return tx, nil
}
func GetTransactionDisplay(id string, tpe string) (*TransactionDisplay, error) {
	tx := &TransactionDisplay{}
	var sql string
	switch tpe {
	case TransTra:
		sql = "Select t.*, t1.name as fromname, t2.name as toname from tx t join account t1 on t.fromid = t1.id join envelope t2 on t.toid = t2.id where id = ?"
	case TransEnv:
		sql = "Select t.*, t1.name as fromname, t2.name as toname from tx t join envelope t1 on t.fromid = t1.id join envelope t2 on t.toid = t2.id where id = ?"
	case TransAcc:
		sql = "Select t.*, t1.name as fromname, t2.name as toname from tx t join account t1 on t.fromid = t1.id join account t2 on t.toid = t2.id where id = ?"
	}

	err := db.Get(tx, sql, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return tx, nil
}
*/
