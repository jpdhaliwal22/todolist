package db

import (
	"github.com/jpdhaliwal22/Todo-Golang/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	CreateUser(model.User) (int64, error)
	GetUser(map[string]interface{}) (model.User, error)
	AddTask(model.Task) (model.Task, error)
	GetTaskList(map[string]interface{}) ([]model.Task, error)
	DeleteTask(taskId string) error
	UpdateTask(model.Task) error
}

type mysqlDb struct {
	db *gorm.DB
}

func NewDB() Database {

	dsn := "root:Pass@123@tcp(localhost:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	db.AutoMigrate(&model.Task{}, &model.User{})

	return &mysqlDb{db: db}
}
