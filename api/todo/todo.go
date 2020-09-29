package todo

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTodo struct {
	Username string `json:"username" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Message  string `json:"message" binding:"required"`
}

func GetAllTodolists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resultGetAll": "ok",
	})
}

func GetTodolists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resultGet": "ok",
	})
}

func PostTodolist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resultPost": "ok",
	})
}

func PutTodolist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resultPut": "ok",
	})
}

func DeleteTodolist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"resultDelete": "ok",
	})
}
