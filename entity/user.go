package entity

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}
