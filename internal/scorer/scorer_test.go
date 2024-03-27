package scorer

import (
	"testing"

	models "superheroes-go/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestHeroWinsWithHigherScore(t *testing.T) {
	hero := models.Character{
		Name:  "Wonder Woman",
		Score: 8.7,
		Type:  "hero",
	}

	villain := models.Character{
		Name:  "Lex Luthor",
		Score: 8,
		Type:  "villain",
	}

	assert.Equal(t, WinnerCalculator(hero, villain), hero)
}

func TestVillainWinsWithHigherScore(t *testing.T) {
	hero := models.Character{
		Name:  "Gamora",
		Score: 8.4,
		Type:  "hero",
	}

	villain := models.Character{
		Name:  "Thanos",
		Score: 9.9,
		Type:  "villain",
	}

	assert.Equal(t, WinnerCalculator(hero, villain), villain)
}

func TestVillainWinsDueToWeakness(t *testing.T) {
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

	assert.Equal(t, WinnerCalculator(hero, villain), villain)
}

func TestHeroWinsAfterWeakness(t *testing.T) {
	hero := models.Character{
		Name:     "Superman",
		Score:    9.6,
		Type:     "hero",
		Weakness: "Lex Luthor",
	}

	villain := models.Character{
		Name:  "Lex Luthor",
		Score: 8,
		Type:  "villain",
	}

	supermanAfterWeakness := models.Character{
		Name:     "Superman",
		Score:    8.6,
		Type:     "hero",
		Weakness: "Lex Luthor",
	}

	assert.Equal(t, WinnerCalculator(hero, villain), supermanAfterWeakness)
}
