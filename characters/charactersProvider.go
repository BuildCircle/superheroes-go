package characters

import (
	"encoding/json"
	"io"
	"net/http"
)

type Character struct {
	Name     string  `json:"name"`
	Score    float32 `json:"score"`
	Type     string  `json:"type"`
	Weakness string  `json:"Weakness"`
}

type Characters struct {
	Items []Character `json:"items"`
}

func GetCharacters() (Characters, error) {
	var characters Characters
	resp, err := http.Get("https://s3.eu-west-2.amazonaws.com/build-circle/characters.json")
	if err != nil {
		return characters, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return characters, err
	}

	jsonErr := json.Unmarshal(body, &characters)

	return characters, jsonErr
}
