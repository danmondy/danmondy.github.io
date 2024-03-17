package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/danmondy/AirshipCards/data"
	"github.com/danmondy/AirshipCards/templates"
	"github.com/labstack/echo/v4"
)

func RegisterPlayHandlers(e *echo.Echo) {
	e.GET("/play/:gameid", GetGame)
}

func GetGame(c echo.Context) error {
	gameID := c.Param("gameid")
	game, err := data.GetById[data.Game](gameID)
	if err != nil {
		return err
	}
	game.Board, err = data.GetById[data.Board](game.BoardID)
	if err != nil {
		return err
	}
	game.Board.Hexes, err = data.GetAllFor[data.Hex](game.BoardID, "boardid")
	game.Players = []data.Player{data.NewPlayer("Dan"), data.NewPlayer("Jon")}

	game.Decks, err = data.GetAll[data.Deck]()
	log.Println(game.Decks)
	for i, d := range game.Decks {
		game.Decks[i].CardList, err = data.GetAllFor[data.Card](d.ID, "deckid")
	}
	log.Println(game.Decks[1].CardList)
	games, err := data.GetAll[data.Game]()
	if err != nil {
		return err
	}
	log.Println("making new game")

	for _, d := range game.Decks {
		ShuffleDeck(d.CardList)
	}

	return render(c, http.StatusOK, templates.GameLayout(games, game))
}

func ShuffleDeck(d []data.Card) {
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
	//return d;
}
