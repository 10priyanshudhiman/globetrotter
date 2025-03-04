package model

type Game struct {
	Clues  map[string]string `json:"clues"`
	Cities []string          `json:"cities"`
}

type VerifyAction struct {
	Country string `json:"country"`
	City    string `json:"city"`
	ClueId  string `json:"clue_id"`
}

type ActionResponse struct {
	IsCorrect bool              `json:"isCorrect"`
	Facts     map[string]string `json:"facts"`
}
