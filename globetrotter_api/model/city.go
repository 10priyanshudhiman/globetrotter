package model

type City struct {
	City    string   `json:"city"`
	Country string   `json:"country"`
	Clues   []string `json:"clues"`
	FunFact []string `json:"fun_fact"`
	Trivia  []string `json:"trivia"`
}
