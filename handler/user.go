package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/jpdhaliwal22/Todo-Golang/entity"
	"github.com/jpdhaliwal22/Todo-Golang/model"
)

var secret = []byte("secret")

func (h *handlerStr) CreateUser(c *gin.Context) {
	fmt.Println("CREATE UJSER")
	var user entity.User

	// Call BindJSON to bind the received JSON to
	// task.
	if err := c.BindJSON(&user); err != nil {
		fmt.Println("ERR", err)
		return
	}

	fmt.Println("user", user)
	m := model.User{
		UserName: user.UserName,
		Password: user.Password,
	}

	fmt.Println("model", m)
	// Add the new album to the slice.
	//albums = append(albums, newAlbum)//
	if user, err := h.svc.CreateUser(m); err != nil {
		c.IndentedJSON(http.StatusBadRequest, user)
		return
	}

	c.IndentedJSON(http.StatusCreated, user)

}
func (h *handlerStr) UserLogin(c *gin.Context) {
	var user entity.User
	session := sessions.Default(c)

	// Call BindJSON to bind the received JSON to
	// task.
	if err := c.BindJSON(&user); err != nil {
		return
	}

	fmt.Println("UserName", user.UserName)
	existingUser, err := h.svc.GetUser(map[string]interface{}{"user_name": user.UserName})

	fmt.Println("existing user", existingUser)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, user)
		return
	}

	if existingUser.Password != user.Password {
		fmt.Println("pass not match", user.Password)
		c.IndentedJSON(http.StatusUnauthorized, user)
		return
	}

	session.Set(userkey, existingUser.ID) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.IndentedJSON(http.StatusAccepted, user)
	c.SetCookie("userid", string(existingUser.ID), 36000, "/", "localhost", false, true)
}
