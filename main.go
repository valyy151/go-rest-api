package main

import (
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Learn Go", Completed: false},
	{ID: "2", Item: "Learn React", Completed: true},
	{ID: "3", Item: "Learn German", Completed: false},
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoByID)

	router.POST("/todos", addTodos)
	router.DELETE("/todos/:id", deleteTodos)
	router.PATCH("/todos/:id", updateTodoByID)

	router.Run("localhost:8080")
}
