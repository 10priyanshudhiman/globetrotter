package model

type Message struct {
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
