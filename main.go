package main

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jpdhaliwal22/Todo-Golang/db"
	"github.com/jpdhaliwal22/Todo-Golang/handler"
	"github.com/jpdhaliwal22/Todo-Golang/service"
)

var secret = []byte("secret")

func main() {
	fmt.Println("Starting")
	r := gin.Default()
	r.Use(sessions.Sessions("mysession", cookie.NewStore(secret)))
	r.Use(CORSMiddleware())
	db := db.NewDB()
	svc := service.NewService(db)
	middleware := handler.NewHandler(svc)
	r.POST("/signup", middleware.CreateUser)
	r.POST("/login", middleware.UserLogin)
	r.POST("/tasks", middleware.AddTask)
	r.GET("/tasks", middleware.GetTaskList)
	r.DELETE("/tasks/:id", middleware.DeleteTask)
	r.PUT("/tasks", middleware.UpdateTask)
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
