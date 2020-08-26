package main

import (
	"api/infra"
	"api/model"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type handler struct {
	databaseHandler infra.DatabaseHandler
}

func newHandler(databaseHandler infra.DatabaseHandler) handler {
	return handler{databaseHandler}
}

func (h *handler) getWords() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := h.databaseHandler.GetWords()
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, data)
	}
}

func (h *handler) updateWord() echo.HandlerFunc {
	return func(c echo.Context) error {
		r := &model.Request{}
		if err := c.Bind(r); err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		err := h.databaseHandler.UpdateWord(r.Word)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusOK)
	}
}

func main() {
	// If you want to use MySQL DB, set MYSQL_CONNECTION in your shell.
	conn := os.Getenv("MYSQL_CONNECTION")
	var databaseHandler infra.DatabaseHandler
	if conn != "" {
		databaseHandler = infra.NewSQLHandler(conn)
	} else {
		databaseHandler = infra.NewMemoryHandler()
	}
	h := newHandler(databaseHandler)
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/words", h.getWords())
	e.POST("/words", h.updateWord())

	e.Start(":1323")
}
