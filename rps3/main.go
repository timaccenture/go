package main

import (
	"math/rand/v2"
	"net/http"
	"os"
	"slices"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var moves = map[string]string{
	"1": "rock",
	"2": "paper",
	"3": "scissor",
}

var userMinusComp = map[string][]int{
	"Client won": {1, 2, -3},
	"Server won": {-1, -2, 3},
}

func playGame(c echo.Context) error {
	id := c.Param("id")
	clientMove := moves[id]
	validInputs := []string{"rock", "paper", "scissor"}
	randNumber := rand.IntN(len(validInputs))
	serverMove := validInputs[randNumber]

	diff := len(clientMove) - len(serverMove)

	if diff == 0 {
		result := "Tied, both chose " + clientMove
		return c.String(http.StatusOK, result)
	} else if slices.Contains(userMinusComp["Client won"], diff) {
		result := "You won with " + clientMove + ", computer chose " + serverMove
		return c.String(http.StatusOK, result)
	} else if slices.Contains(userMinusComp["Server won"], diff) {
		result := "Computer won with " + serverMove + ", you chose " + clientMove
		return c.String(http.StatusOK, result)
	} else {
		return c.String(http.StatusOK, "Ooops, sth went wrong!")
	}

}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/game", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Play rock paper scissors against the computer...")
	})

	e.POST("/game/:id", playGame)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
