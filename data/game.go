package data

import (
	"image/color"

	"github.com/danmondy/AirshipCards/data"
)

type Game struct {
	ShopDeck   []Card
	EventDeck  []Card
	ActionDeck []Card
	Players    []Player
}
func NewGame()Game{
	cards, err := data.GetDeckByName[Card]("shop")
	if err != nil{
		log.Fatal(err)
	}
	return Game{
		ShopDeck: data
}
type Hex struct {
	Discovered bool
	Type       string
}
type Player struct {
	Name     string `db:"name" form:"text" display:"Name"`
	Color    color.Color
	Deck     []Card
	PortBank Bank
	ShipBank Bank
}
type Bank struct {
	Wood  int
	Ore   int
	Adder int
	Wool  int
}
