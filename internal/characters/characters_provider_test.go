package characters

import (
	"log"
	models "superheroes-go/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetsCharacters(t *testing.T) {
	batman := models.Character{
		Name:     "Batman",
		Score:    8.3,
		Type:     "hero",
		Weakness: "Joker",
	}

	joker := models.Character{
		Name:  "Joker",
		Score: 8.2,
		Type:  "villain",
	}

	hero, villain, err := GetCharacters("Batman", "Joker")
	if err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, hero, batman)
	assert.Equal(t, villain, joker)
}
