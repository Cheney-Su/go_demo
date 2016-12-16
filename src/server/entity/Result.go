package entity

type Result struct {
	Status int	`json:"status"`
	Data interface{}
	Msg string	`json:"msg"`
}