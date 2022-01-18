package entity

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

//Model / Struct Response
type ResponseUser struct {
	Http    string     `json:"http"`
	Message string     `json:"message"`
	Data    Tb_account `json:"data"`
}

type ResponseUsers struct {
	Http    string       `json:"http"`
	Message string       `json:"message"`
	Data    []Tb_account `json:"data"`
}

type Tb_account struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
}
