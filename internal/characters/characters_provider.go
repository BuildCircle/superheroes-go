package characters

import (
	"encoding/json"
	"io"
	"net/http"
	models "superheroes-go/internal/models"
)

type Characters struct {
	Items []models.Character `json:"items"`
}

func GetCharacters(hero string, villain string) (models.Character, models.Character, error) {
	var heroChar models.Character
	var villainChar models.Character
	var characters Characters
	resp, err := http.Get("https://s3.eu-west-2.amazonaws.com/build-circle/characters.json")
	if err != nil {
		return heroChar, villainChar, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return heroChar, villainChar, err
	}

	err = json.Unmarshal(body, &characters)
	if err != nil {
		return heroChar, villainChar, err
	}

	for _, item := range characters.Items {
		if item.Name == hero {
			heroChar = item
		}
		if item.Name == villain {
			villainChar = item
		}
	}

	return heroChar, villainChar, nil
}
