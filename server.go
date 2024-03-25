package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"superheroes-go/characters"
)

type Character = characters.Character

func main() {
	e := echo.New()
	var heroChar Character
	var villainChar Character
	
	e.GET("/battle", func(c echo.Context) error {

		hero := c.QueryParam("hero")
		villain := c.QueryParam("villain")

		characters, err := characters.GetCharacters()

		if err != nil {
			e.Logger.Fatal(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		for _, item := range characters.Items {
			if item.Name == hero {
				heroChar = item
			}
			if item.Name == villain {
				villainChar = item
			}
		}

		if heroChar.Score > villainChar.Score {
			return c.JSON(http.StatusOK, heroChar)
		}

		return c.JSON(http.StatusOK, villainChar)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
