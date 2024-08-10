package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jpdhaliwal22/Todo-Golang/entity"
	"github.com/jpdhaliwal22/Todo-Golang/model"
)

const userkey = "user"

func (h *handlerStr) AddTask(c *gin.Context) {
	var task entity.Task

	session := sessions.Default(c)
	userID := session.Get(userkey)
	// Call BindJSON to bind the received JSON to
	// task.
	if err := c.BindJSON(&task); err != nil {
		return
	}

	taskM := model.Task{
		Status: "InProgress",
		Detail: task.Detail,
		UserID: userID.(uint),
	}

	fmt.Println("userID", taskM.UserID)
	// Add the new album to the slice.
	//albums = append(albums, newAlbum)//
	rTask, err := h.svc.AddTask(taskM)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, task)
		return
	}

	c.IndentedJSON(http.StatusCreated, rTask)

}

func (h *handlerStr) GetTaskList(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get(userkey)

	tasks, err := h.svc.GetTaskList(map[string]interface{}{"user_id": userID})
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, tasks)
		return
	}

	c.IndentedJSON(http.StatusOK, tasks)
}

func (h *handlerStr) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")

	err := h.svc.DeleteTask(taskID)
	if err != nil {
		fmt.Println("ERROR task", taskID)
		c.IndentedJSON(http.StatusNotFound, taskID)
		return
	}

	c.IndentedJSON(http.StatusOK, taskID)

}

func (h *handlerStr) UpdateTask(c *gin.Context) {
	var task entity.Task
	session := sessions.Default(c)
	userID := session.Get(userkey)

	if err := c.BindJSON(&task); err != nil {
		return
	}

	m := model.Task{
		Status: task.Status,
		Detail: task.Detail,
		UserID: userID.(uint),
		ID:     task.ID,
	}

	err := h.svc.UpdateTask(m)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, m)
		return
	}

	c.IndentedJSON(http.StatusCreated, m)

}
