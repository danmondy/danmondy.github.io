package data

import (
	"image/color"
	"math/rand/v2"
)

type Game struct {
	ID      string `db:"id,primarykey" form-type:"hidden" form:"id"`
	BoardID string `db:"boardid" form-type:"text" form:"boardid" label:"board"` //todo: make select list
	Name    string `db:"name" form-type:"text" form:"name" label:"name"`
	Decks   []Deck `db:"omit" form-type:"omit"`
	//ShopDeck   []Card   `db:"omit"`
	//EventDeck  []Card   `db:"omit"`
	//ActionDeck []Card   `db:"omit"`
	Players []Player `db:"omit" form-type:"omit"`
	Board   Board    `db:"omit" form-type:"omit"`
}

func NewGame() Game {
	//GetDecksForGame()
	return Game{
		ID:      NewUniqueID(),
		Name:    "new game",
		BoardID: "",
	}
}

type Player struct {
	Name     string `db:"name" form:"text" display:"Name"`
	Color    string
	Deck     map[int]*Card
	PortBank Bank
	ShipBank Bank
}

func NewPlayer(name string) Player {
	return Player{
		Name:  name,
		Color: GetHexFromColor(color.RGBA{R: uint8(rand.IntN(155)), G: uint8(rand.IntN(155)), B: uint8(rand.IntN(155)), A: 255}),
		Deck:  map[int]*Card{0: nil, 1: nil, 2: nil, 3: nil, 4: nil, 5: nil},
		//PortBank: Bank{},
		//ShipBank: Bank{},
	}
}

type Bank struct {
	Wood  int
	Ore   int
	Adder int
	Wool  int
}
