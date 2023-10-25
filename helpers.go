package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodo(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found")
}

func getTodoByID(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodo(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func addTodos(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func updateTodo(id string, updatedTodo todo) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			todos[i] = updatedTodo
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found")
}

func updateTodoByID(context *gin.Context) {
	id := context.Param("id")

	var updatedTodo todo

	if err := context.BindJSON(&updatedTodo); err != nil {
		return
	}

	todo, err := updateTodo(id, updatedTodo)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func deleteTodos(context *gin.Context) {
	id := context.Param("id")

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}

	context.IndentedJSON(http.StatusOK, todos)
}
