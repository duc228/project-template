package app

import (
	"server/internal/common/error"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewResponse(c *gin.Context, errCode int, data interface{}) {
	if data == nil {
		data = error.GetMsg(errCode)
	}

	c.JSON(errCode, Response{
		Code: errCode,
		Data: data,
	})
}
