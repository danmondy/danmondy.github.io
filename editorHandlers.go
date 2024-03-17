package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/danmondy/AirshipCards/data"
	"github.com/danmondy/AirshipCards/templates"
	"github.com/labstack/echo/v4"
)

func RegisterEditorHandlers(e *echo.Echo) {
	e.GET("/editor/:gameid", GetEditor)

	e.GET("/editor/game/new", GetGameNew)
	e.POST("/editor/game/new", PostGameNew)
	e.GET("/editor/game/edit/:id", GameEdit)
	e.PUT("/editor/game/update", GameUpdate)
	e.DELETE("/editor/game/:id", GameDelete)

	e.GET("/editor/decks", GetAllDecks)
	e.GET("/editor/deck/:id", GetDeck)
	e.GET("/editor/deck/edit/:id", DeckEdit)
	e.DELETE("/editor/deck/:id", DeckDelete)
	e.PUT("/editor/deck/update", DeckUpdate)
	e.POST("/editor/deck/new", DeckNew)

	e.GET("/editor/card/edit/:id", CardEdit)
	e.PUT("/editor/card/update", CardUpdate)
	e.POST("/editor/card/:deckid", CardNew)
	e.DELETE("/editor/card/:id", CardDelete)

	e.GET("/editor/board/:id", GetBoard)
	e.GET("/editor/board/edit/:id", BoardEdit)
	e.PUT("/editor/board/update", BoardUpdate)
	e.DELETE("/editor/board/:id", BoardDelete)

	e.GET("/editor/hex/edit/:id", HexEdit)
	e.PUT("/editor/hex/update", HexUpdate)
	e.POST("/editor/hex/new/:boardid", PostHexNewQuick)
	e.POST("/editor/hex/:boardid", PostHexNew)
	e.DELETE("/editor/hex/:id", HexDelete)
}

func GetEditor(c echo.Context) error {
	gameID := c.Param("gameid")
	var game data.Game
	games, err := data.GetAll[data.Game]()
	if err != nil {
		return err
	}
	for _, g := range games {
		if g.ID == gameID {
			game = g
		}
	}
	decks, err := data.GetAll[data.Deck]()
	if err != nil {
		return err
	}
	game.Decks = decks
	var board data.Board
	if game.BoardID != "" {
		board, err = data.GetById[data.Board](game.BoardID)
		if err != nil {
			board = data.Board{ID: game.BoardID, Name: "Game Board", Hexes: make([]data.Hex, 0)}
			data.Insert(&board)
		}
	}
	game.Board = board //this board doesn't have hexes yet, it's just the name and id
	return render(c, http.StatusOK, templates.EditorLayout(games, game))
}

func GetGameNew(c echo.Context) error {
	game := data.NewGame()
	return render(c, 200, templates.CreateForm(&game, true))
}
func PostGameNew(c echo.Context) error {
	game := &data.Game{}
	err := c.Bind(game)
	if err != nil {
		return err
	}
	board := data.NewBoard()
	data.Insert(&board)
	game.BoardID = board.ID
	data.Insert(game)

	c.Response().Header().Set("HX-Redirect", "/editor/"+game.ID)
	return c.String(200, "Saved")
}
func GameEdit(c echo.Context) error {
	gameID := c.Param("id")

	game, err := data.GetById[data.Game](gameID)
	if err != nil {
		return err
	}
	return render(c, http.StatusOK, templates.CreateForm(&game, false))
}
func GameUpdate(c echo.Context) error {
	var game data.Game

	err := c.Bind(&game)
	if err != nil {
		return err
	}
	data.Update(&game)

	c.Response().Header().Set("HX-Trigger", "RefreshGame")
	return c.String(http.StatusOK, "")
}
func GameDelete(c echo.Context) error {
	fmt.Println("Delete Reached")
	id := c.Param("id")
	g, err := data.GetById[data.Game](id)
	if err != nil {
		return err
	}
	for _, gd := range g.Decks {
		data.Delete("gamedeck", gd.ID)
	}
	data.Delete("game", id)

	c.Response().Header().Set("HX-Redirect", "/")
	return c.String(http.StatusOK, "")
}

func DeckEdit(c echo.Context) error {
	deckID := c.Param("id")

	deck, err := data.GetById[data.Deck](deckID)
	if err != nil {
		return err
	}
	return render(c, http.StatusOK, templates.CreateForm(&deck, false))
}
func DeckUpdate(c echo.Context) error {
	deck := data.Deck{
		ID:   c.FormValue("ID"),
		Name: c.FormValue("Name"),
	}
	data.Update(&deck)

	c.Response().Header().Set("HX-Trigger", "RefreshDecks")
	return c.String(http.StatusOK, "")
}

func DeckDelete(c echo.Context) error {
	id := c.Param("id")
	data.DeleteCardsInDeck(id)
	data.DeleteGameDeckByDeck(id)
	data.Delete("deck", id)

	//data.DeleteDeckCard()
	//todo: get rid of the deckcard table - thats dumb

	c.Response().Header().Set("HX-Trigger", "RefreshDecks")
	return c.String(http.StatusOK, "deleted")
}
func GetAllDecks(c echo.Context) error {
	decks, err := data.GetAll[data.Deck]()
	if err != nil {
		return err
	}
	return render(c, http.StatusOK, templates.DeckList(decks))
}

func GetDeck(c echo.Context) error {
	deckID := c.Param("id")
	deck, err := data.GetById[data.Deck](deckID)
	if err != nil {
		log.Println(err)
		return err
	}

	cards, err := data.GetAllForSorted[data.Card](deck.ID, "deckid", "text")
	if err != nil {
		log.Println(err)
	}
	return render(c, http.StatusOK, templates.DeckView(deck, cards))
}

func DeckNew(c echo.Context) error {
	deck := data.NewDeck()
	data.Insert(&deck)

	c.Response().Header().Set("HX-Trigger", "RefreshDecks")
	return c.String(200, "")
}

func CardNew(c echo.Context) error {
	deckID := c.Param("deckid")
	card := data.NewCard(deckID)

	data.Insert(&card)

	return render(c, http.StatusOK, templates.Card(card))
}

func CardEdit(c echo.Context) error {
	cardID := c.Param("id")

	card, err := data.GetById[data.Card](cardID)
	if err != nil {
		return err
	}
	return render(c, http.StatusOK, templates.CreateForm(&card, false))
}
func CardUpdate(c echo.Context) error {

	var card data.Card

	err := c.Bind(&card)
	if err != nil {
		return err
	}
	data.Update(&card)

	c.Response().Header().Set("HX-Trigger", "RefreshDeck")
	return c.String(http.StatusOK, "")
}

func CardDelete(c echo.Context) error {
	fmt.Println("Delete Reached")
	id := c.Param("id")
	data.Delete("card", id)
	//data.DeleteDeckCard()
	//todo: get rid of the deckcard table - thats dumb

	c.Response().Header().Set("HX-Trigger", "RefreshDeck")
	return c.String(http.StatusOK, "")
}

func BoardReset(c echo.Context) error {
	//todo clear all the hexes
	BoardID := c.Param("id")
	err := data.DeleteAll[data.Hex](BoardID, "boardid")
	if err != nil {
		return err
	}
	return nil
}

func GetBoard(c echo.Context) error {
	BoardID := c.Param("id")

	board, err := data.GetById[data.Board](BoardID)
	if err != nil {
		return err
	}
	board.Hexes, err = data.GetAllFor[data.Hex](BoardID, "boardid")
	if err != nil {
		return err
	}

	template := data.NewBoardTemplate(25, 16)

	return render(c, http.StatusOK, templates.BoardEditor(board, template))
}

func BoardEdit(c echo.Context) error {
	BoardID := c.Param("id")

	board, err := data.GetById[data.Board](BoardID)
	if err != nil {
		return err
	}

	return render(c, http.StatusOK, templates.CreateForm(&board, false))
}
func BoardUpdate(c echo.Context) error {

	var Board data.Board

	err := c.Bind(&Board)
	if err != nil {
		return err
	}
	data.Update(&Board)

	c.Response().Header().Set("HX-Trigger", "RefreshBoard")
	return c.String(http.StatusOK, "")
}

func BoardDelete(c echo.Context) error {
	fmt.Println("Delete Reached")
	id := c.Param("id")
	data.Delete("Board", id)
	data.DeleteAll[data.Hex](id, "boardid")
	//todo: get rid of the deckBoard table - thats dumb

	c.Response().Header().Set("HX-Trigger", "RefreshDeck")
	return c.String(http.StatusOK, "")
}

/*
func GetHexNew(c echo.Context) error {
	boardID := c.Param("boardid")
	x, err := strconv.Atoi(c.QueryParam("x"))
	if err != nil {
		return err
	}
	y, err := strconv.Atoi(c.QueryParam("y"))
	if err != nil {
		return err
	}
	fmt.Println(boardID, x, y)
	Hex := data.Hex{ID: data.NewUniqueID(), BoardID: boardID, X: x, Y: y}

	data.Insert(&Hex)

	return render(c, http.StatusOK, templates.CreateForm(&Hex, true))
}*/

func PostHexNewQuick(c echo.Context) error {
	ac := c.FormValue("activeColor")

	boardID := c.Param("boardid")
	x, err := strconv.Atoi(c.QueryParam("x"))
	if err != nil {
		return err
	}
	y, err := strconv.Atoi(c.QueryParam("y"))
	if err != nil {
		return err
	}
	h := data.Hex{ID: data.NewUniqueID(), BoardID: boardID, Color: ac, X: x, Y: y, Type: ""}
	data.Insert(&h)

	c.Response().Header().Set("HX-Trigger", "RefreshBoard")
	return c.String(200, "Saved")
}

func PostHexNew(c echo.Context) error {
	h := &data.Hex{}
	err := c.Bind(h)
	if err != nil {
		return err
	}
	data.Insert(h)

	c.Response().Header().Set("HX-Trigger", "RefreshBoard")
	return c.String(200, "Saved")
}

func HexEdit(c echo.Context) error {
	HexID := c.Param("id")

	color := c.QueryParam("activeColor")
	log.Println("color: ", color)
	hex, err := data.GetById[data.Hex](HexID)
	if err != nil {
		return err
	}
	if color != "" {
		hex.Color = color
	}
	return render(c, http.StatusOK, templates.CreateForm(&hex, false))
}
func HexUpdate(c echo.Context) error {
	var Hex data.Hex
	err := c.Bind(&Hex)
	if err != nil {
		return err
	}
	data.Update(&Hex)

	c.Response().Header().Set("HX-Trigger", "RefreshBoard")
	return c.String(http.StatusOK, "")
}

func HexDelete(c echo.Context) error {
	fmt.Println("Delete Reached")
	id := c.Param("id")
	data.Delete("Hex", id)
	//data.DeleteDeckHex()
	//todo: get rid of the deckHex table - thats dumb

	c.Response().Header().Set("HX-Trigger", "RefreshBoard")
	return c.String(http.StatusOK, "")
}
