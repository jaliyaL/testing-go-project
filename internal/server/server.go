package server

import (
	"myapp/internal/domain/todo"

	"github.com/gin-gonic/gin"
)

func NewServer(todoHandler *todo.Handler) *gin.Engine {
	r := gin.Default()

	t := r.Group("/todos")
	{
		t.POST("/", todoHandler.CreateTodo)
		t.GET("/:id", todoHandler.GetTodoByID)
	}

	return r
}
