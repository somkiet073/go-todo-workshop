package main

import (
	"go-todo-workshop/api/todo"
	// "go-todo-workshop/database"
	// "fmt"
	"net/http" // add import net/http

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin" // add import github.com/gin-gonic/gin
)

func main() {
	// fmt.Println("Hello main")

	r := gin.Default() // add gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://locahost:8080"}
	r.Use(cors.New(config))

	// database

	// step4 Create method GET
	r.GET("/", todo.GetAllTodolists)
	r.GET("/gettodo", todo.GetTodolists)
	r.POST("/", todo.PostTodolist)
	r.PUT("/:id", todo.PutTodolist)
	r.DELETE("/:id", todo.DeleteTodolist)
	r.StaticFS("/file", http.Dir("public"))
	// Final Step Run()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
