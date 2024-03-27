package main

import (
	"net/http"
	provider "superheroes-go/internal/characters"
	"superheroes-go/internal/scorer"
	"superheroes-go/internal/validator"

	"github.com/labstack/echo/v4"
)

type JsonResponse struct {
	Status  int
	Message string
}

func main() {
	e := echo.New()

	e.GET("/battle", func(c echo.Context) error {
		hero := c.QueryParam("hero")
		villain := c.QueryParam("villain")

		if hero == "" || villain == "" {
			jsonResponse := JsonResponse{
				Status:  http.StatusBadRequest,
				Message: "Missing one or more arguments",
			}
			return c.JSON(http.StatusBadRequest, jsonResponse)
		}

		heroChar, villainChar, err := provider.GetCharacters(hero, villain)

		if err != nil {
			e.Logger.Fatal(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		validCharacters := validator.ValidCombatantCharacters(heroChar, villainChar)
		if !validCharacters {
			jsonResponse := JsonResponse{
				Status:  http.StatusUnprocessableEntity,
				Message: "Characters must be of the correct type",
			}
			return c.JSON(http.StatusUnprocessableEntity, jsonResponse)
		}

		winner := scorer.WinnerCalculator(heroChar, villainChar)

		return c.JSON(http.StatusOK, winner)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
