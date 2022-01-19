package entity

type File struct {
	Id   int64  `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name" validate:"required"`
	Path string `gorm:"not null" json:"path" validate:"required"`
}
