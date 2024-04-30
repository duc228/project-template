package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	httpRequest "server/internal/common/utils/http_request"
	"server/pkg/app"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var requestHeader map[string][]string
	var todoList []map[string]interface{}

	resp, e := httpRequest.Get("http://jsonplaceholder.typicode.com/todos", requestHeader, nil)
	if e != nil {
		fmt.Printf("client: could not read response body: %s\n", e)
		app.NewResponse(c, http.StatusBadRequest, nil)
	}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		app.NewResponse(c, http.StatusBadRequest, nil)

	}

	if err := json.Unmarshal([]byte(resBody), &todoList); err != nil {
		app.NewResponse(c, http.StatusBadRequest, nil)
	}

	app.NewResponse(c, http.StatusOK, todoList)
}
