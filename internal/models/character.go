package models

type Character struct {
	Name     string  `json:"name"`
	Score    float32 `json:"score"`
	Type     string  `json:"type"`
	Weakness string  `json:"Weakness,omitempty"`
}
