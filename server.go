package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	files := http.FileServer(rice.MustFindBox("public").HTTPBox())

	e.GET("/", index)
	// urlから/staticを削除して、残りのファイル名をpublic配下から返却
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", files)))

	e.Logger.Fatal(e.Start(":8080"))
}

func index(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}
