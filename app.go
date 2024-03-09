package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/danmondy/AirshipCards/data"
	"github.com/danmondy/AirshipCards/templates"
)

var cfg *data.Config
var EventDeck *data.Deck
var ShopDeck *data.Deck

func init() {
	cfg = &data.Config{}
	err := data.ReadJSON("config.json", cfg)
	if err != nil {
		log.Println(err)
	}

	data.Initialize(false)

	//data.RunSql()
	//data.ReadJSON("deck/event.json", EventDeck)
	//data.ReadJSON("deck/shop.json", ShopDeck)
}

func main() {
	e := echo.New()

	e.Static("/assets", "assets")
	e.GET("/", GetIndex)

	e.GET("/decks", GetAllDecks)
	e.GET("/deck/:id", GetDeck)
	e.GET("/deck/edit/:id", DeckEdit)
	e.DELETE("/deck/:id", DeckDelete)
	e.PUT("/deck/update", DeckUpdate)
	e.POST("/deck/new", DeckNew)

	e.GET("/card/edit/:id", CardEdit)
	e.PUT("/card/update", CardUpdate)
	e.POST("/card/:deckid", CardNew)
	e.DELETE("/card/:id", CardDelete)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}

// this is used by ever handler to render a view with Echo and Templ
func render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func GetIndex(c echo.Context) error {
	decks, err := data.GetAll[data.Deck]()
	if err != nil {
		log.Println(err)
	}
	if len(decks) < 1 {
		return c.String(http.StatusOK, "No Decks")
	}
	cards, err := data.GetAllFor[data.Card](decks[1].ID, "deckid")
	if err != nil {
		log.Println(err)
	}
	return render(c, http.StatusOK, templates.Layout(decks, cards))
}
func DeckEdit(c echo.Context) error {
	deckID := c.Param("id")

	deck, err := data.GetById[data.Deck](deckID)
	if err != nil {
		return err
	}
	return render(c, http.StatusOK, templates.CreateForm(&deck))
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
	data.Delete("deck", id)

	//data.DeleteDeckCard()
	//todo: get rid of the deckcard table - thats dumb

	c.Response().Header().Set("HX-Trigger", "RefreshDecks")
	return c.String(http.StatusOK, "")
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
	log.Println(cards)
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
	return render(c, http.StatusOK, templates.CreateForm(&card))
}
func CardUpdate(c echo.Context) error {
	tr, err := strconv.Atoi(c.FormValue("TextTopRight"))
	if err != nil {
		return err
	}
	card := data.Card{
		ID:             c.FormValue("ID"),
		DeckID:         c.FormValue("DeckID"),
		Text:           c.FormValue("Text"),
		Image:          c.FormValue("Image"),
		TextTopRight:   tr,
		TextBottomLeft: c.FormValue("TextBottomLeft"),
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
