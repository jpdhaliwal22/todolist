package entity

type Task struct {
	ID     uint
	Detail string `Json:"detail"`
	Status string `Json:"status"`
	UserID uint   `Json:"user_id"`
}
