package validator

import (
	models "superheroes-go/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeroCanFightVillain(t *testing.T) {
	hero := models.Character{
		Name:     "Batman",
		Score:    8.3,
		Type:     "hero",
		Weakness: "Joker",
	}

	villain := models.Character{
		Name:  "Joker",
		Score: 8.2,
		Type:  "villain",
	}

	assert.True(t, ValidCombatantCharacters(hero, villain))
}

func TestHeroCantFightHero(t *testing.T) {
	hero := models.Character{
		Name:     "Batman",
		Score:    8.3,
		Type:     "hero",
		Weakness: "Joker",
	}

	assert.False(t, ValidCombatantCharacters(hero, hero))
}

func TestVillainCantFightVillain(t *testing.T) {
	villain := models.Character{
		Name:  "Joker",
		Score: 8.2,
		Type:  "villain",
	}

	assert.False(t, ValidCombatantCharacters(villain, villain))
}

func TestCantBeSwtichedmodels(t *testing.T) {
	hero := models.Character{
		Name:     "Batman",
		Score:    8.3,
		Type:     "hero",
		Weakness: "Joker",
	}

	villain := models.Character{
		Name:  "Joker",
		Score: 8.2,
		Type:  "villain",
	}

	assert.False(t, ValidCombatantCharacters(villain, hero))
}
