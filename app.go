package main

import (
	"fmt"
	"log"
	"net/http"

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

	RegisterPlayHandlers(e)
	RegisterEditorHandlers(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}

// this is used by ever handler to render a view with Echo and Templ
func render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func GetIndex(c echo.Context) error {
	games, err := data.GetAll[data.Game]()
	if err != nil {
		return err
	}

	return render(c, http.StatusOK, templates.Layout(games))
}
