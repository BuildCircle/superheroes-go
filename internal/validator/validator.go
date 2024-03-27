package validator

import (
	models "superheroes-go/internal/models"
)

func ValidCombatantCharacters(hero models.Character, villain models.Character) bool {
	if hero.Type != "hero" || villain.Type != "villain" {
		return false
	}

	return true
}
