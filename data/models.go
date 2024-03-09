package data

import "slices"

type Card struct {
	ID             string `db:"id,primarykey" json:"id" form:"hidden"`
	DeckID         string `db:"deckid" form:"hidden"`
	Text           string `db:"text" json:"text" form:"text" label:"Text"`
	Image          string `db:"image" form:"text" label:"Image"`
	TextTopRight   int    `db:"texttopright" json:"texttopright" form:"number" label:"Top Right"`
	TextBottomLeft string `db:"textbottomleft" json:"textbottomleft" form:"text" label:"Bottom Left"`
}

func NewCard(deckID string) Card {
	return Card{ID: NewUniqueID(), DeckID: deckID, Text: "", TextTopRight: 1, TextBottomLeft: ""}
}

type Deck struct {
	ID   string `db:"id,primarykey" json:"id" form:"hidden"`
	Name string `db:"name" json:"name" form:"text" label:"Name"`
}

func NewDeck() Deck {
	return Deck{ID: NewUniqueID(), Name: "New Deck"}
}

type CardDeck []Card

func (d CardDeck) Draw() Card {
	c := d[len(d)-1]
	d = slices.Delete(d, len(d)-1, len(d))

	return c
}
