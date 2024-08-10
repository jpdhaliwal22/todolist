package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jpdhaliwal22/Todo-Golang/service"
)

type Handler interface {
	CreateUser(c *gin.Context)
	UserLogin(c *gin.Context)
	AddTask(c *gin.Context)
	GetTaskList(c *gin.Context)
	DeleteTask(c *gin.Context)
	UpdateTask(c *gin.Context)
}

type handlerStr struct {
	svc service.Service
}

func NewHandler(s service.Service) Handler {
	return &handlerStr{svc: s}
}
