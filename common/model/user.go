package model

type User struct {
	UserId           string `json:"userId"`
	TotalScore       int64  `json:"totalScore"`
	CorrectAnswers   int64  `json:"correctAnswers"`
	IncorrectAnswers int64  `json:"incorrectAnswers"`
}

type UserId struct {
	UserId string `json:"userId"`
}

type UserAction struct {
	UserId    string `json:"userId"`
	Point     int64  `json:"point"`
	IsCorrect bool   `json:"isCorrect"`
}
