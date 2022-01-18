package entity

type People struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name" validate:"required,alpha,len=20"`
	Last_name  string `json:"last_name" validate:"required,alpha,len=20"`
	Age        int    `json:"age" validate:"required,numeric,gte=20,lte=65"`
}
