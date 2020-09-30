package todo

import (
	"go-todo-workshop/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTodolists(c *gin.Context) {
	db := database.ConnectDatabase()

	selDB, err := db.Prepare("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()

	rows, err := selDB.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var dataRespone []database.User
	for rows.Next() {
		var users database.User
		rows.Scan(&users.ID, &users.Username, &users.Password, &users.Fullname)
		dataRespone = append(dataRespone, users)
	}
	c.JSON(http.StatusOK, gin.H{
		"resultGetAll": dataRespone,
	})
}

func GetTodolists(c *gin.Context) {
	db := database.ConnectDatabase()
	nID := c.Query("id")

	selDB, err := db.Query("SELECT * FROM users WHERE id=?", nID)
	if err != nil {
		panic(err.Error())
	}
	defer selDB.Close()

	var dataRespone []database.User
	for selDB.Next() {
		var users database.User
		selDB.Scan(&users.ID, &users.Username, &users.Password, &users.Fullname)
		dataRespone = append(dataRespone, users)
	}
	c.JSON(http.StatusOK, gin.H{
		"resultGet": dataRespone,
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
