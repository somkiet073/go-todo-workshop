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
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}
	defer selDB.Close()

	rows, err := selDB.Query()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
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

	db := database.ConnectDatabase()
	var input database.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	insForm, err := db.Prepare("INSERT INTO users(username, password, fullname) VALUES(?,?,?)")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	todoList := database.User{
		Username: input.Username,
		Password: input.Password,
		Fullname: input.Fullname,
	}

	insForm.Exec(input.Username, input.Password, input.Fullname)

	c.JSON(http.StatusOK, gin.H{
		"resultPost": todoList,
	})
}

func PutTodolist(c *gin.Context) {
	id := c.Param("id")
	db := database.ConnectDatabase()
	var input database.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	insForm, err := db.Prepare("UPDATE users SET username=?, password=?, fullname=? WHERE id=?")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}
	insForm.Exec(input.Username, input.Password, input.Fullname, id)
	c.JSON(http.StatusOK, gin.H{
		"resultPut": "ok",
	})
}

func DeleteTodolist(c *gin.Context) {
	db := database.ConnectDatabase()
	id := c.Param("id")

	del, err := db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}
	del.Exec(id)
	c.JSON(http.StatusOK, gin.H{
		"resultDelete": "ok",
	})
}
