package entity

type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}
