package entity

type Book struct {
	Id          int64  `json:"id"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Pages       int    `json:"pages"`
	Price       int    `json:"price"`
	Publication string `json:"publication"`
}
