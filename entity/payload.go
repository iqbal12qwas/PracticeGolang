package entity

type Response struct {
	Http    string      `json:"http"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
