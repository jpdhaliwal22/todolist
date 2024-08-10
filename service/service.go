package service

import (
	"github.com/jpdhaliwal22/Todo-Golang/db"
	"github.com/jpdhaliwal22/Todo-Golang/model"
)

type Service interface {
	CreateUser(model.User) (int, error)
	GetUser(map[string]interface{}) (model.User, error)
	AddTask(model.Task) (model.Task, error)
	GetTaskList(map[string]interface{}) ([]model.Task, error)
	DeleteTask(taskId string) error
	UpdateTask(model.Task) error
}

type svc struct {
	db db.Database
}

func NewService(d db.Database) Service {
	return &svc{db: d}
}
