package data

type Card struct {
	ID             string `db:"id,primarykey" form:"id" json:"id" form-type:"hidden"`
	DeckID         string `db:"deckid" form:"deckid" form-type:"hidden"`
	Text           string `db:"text" form:"text" json:"text" form-type:"text" label:"Text"`
	Image          string `db:"image" form:"image" form-type:"text" label:"Image"`
	TextTopRight   int    `db:"texttopright" form:"textopright" json:"texttopright" form-type:"number" label:"Top Right"`
	TextBottomLeft string `db:"textbottomleft" form:"textbottomleft" json:"textbottomleft" form-type:"text" label:"Bottom Left"`
	Wood           string `db:"wood" form:"wood" form-type:"number" label:"wood"`
	Ore            string `db:"ore" form:"ore" form-type:"number" label:"ore"`
	Adder          string `db:"adder" form:"adder" form-type:"number" label:"adder"`
	Fiber          string `db:"fiber" form:"fiber" form-type:"number" label:"fiber"`
}

func NewCard(deckID string) Card {
	return Card{ID: NewUniqueID(), DeckID: deckID, Text: "", TextTopRight: 1, TextBottomLeft: ""}
}

type CardList []Card

type Deck struct {
	CardList `db:"omit" form-type:"omit"`
	ID       string `db:"id,primarykey" json:"id" form-type:"hidden"`
	Name     string `db:"name" json:"name" form-type:"text" label:"Name"`
}

func NewDeck() Deck {
	return Deck{ID: NewUniqueID(), Name: "New Deck"}
}
