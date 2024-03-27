package scorer

import (
	models "superheroes-go/internal/models"
)

func WinnerCalculator(hero models.Character, villain models.Character) models.Character {
	if hero.Weakness == villain.Name {
		hero.Score -= 1
	}

	if hero.Score > villain.Score {
		return hero
	}

	return villain
}
