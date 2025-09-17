package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func getUsers(c *gin.Context) {
	users := []User{
		{ID: 1, Username: "Supachok"},
	}
	c.JSON(200, users)
}

func main() {
	r := gin.Default()

	r.GET("/users", getUsers)

	r.Run(":8080")
}
