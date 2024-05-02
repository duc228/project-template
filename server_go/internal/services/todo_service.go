package services

import (
	"encoding/json"
	"io"

	httpRequest "server/internal/common/utils/http_request"
)

type TodoService struct {
}

func (s *TodoService) GetTodos() (interface{}, error) {
	var requestHeader map[string][]string
	var todoList interface{}

	resp, err := httpRequest.Get("http://pt_py:5001/api/todo", requestHeader, nil)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(resBody), &todoList); err != nil {
		return nil, err
	}

	return todoList, nil
}
