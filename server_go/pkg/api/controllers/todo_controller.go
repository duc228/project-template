package controllers

import (
	"net/http"
	s "server/internal/services"
	"server/pkg/app"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todoService = s.TodoService{}
	var todoList interface{}

	todoList, err := todoService.GetTodos()
	if err != nil {
		app.NewResponse(c, http.StatusInternalServerError, nil)
		return
	}

	app.NewResponse(c, http.StatusOK, todoList)
}
